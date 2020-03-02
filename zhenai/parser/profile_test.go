package parser

import (
	"io/ioutil"
	"testing"
	"webcrawler/engine"
	"webcrawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.txt")
	if err != nil {
		t.Errorf("an error occured when opening test file ./profile_test_data.txt")
		return
	}

	result := ParseProfile(contents, "你值得拥有", "http://album.zhenai.com/1234567")
	if len(result.Items) != 1 {
		t.Errorf(
			"result.Items should contain 1 element, but had %d elements", len(result.Items))
	}
	actual := result.Items[0]

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
	if actual != expected {
		t.Errorf("profiles not match")
	}
}
