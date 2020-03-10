package persist

import (
	"context"
	"encoding/json"
	"testing"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/model"

	"github.com/olivere/elastic/v7"
)

const index = "dating_test"

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/1234567",
		Id:  "1234567",
		Payload: model.Profile{
			Name:     "你值得拥有",
			Age:      24,
			Height:   168,
			Location: "阿坝九寨沟",
			Income:   "3千以下",
		},
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, err := model.FromJsonObjToProfile(actual.Payload)
	if err != nil {
		panic(err)
	}
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %+v, but expected %+v", actual, expected)
	}
}
