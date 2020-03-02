package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	const resultSize = 470
	const itemSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	contents, err := ioutil.ReadFile("./citylist_test_data.txt")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents, "")

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != itemSize {
		t.Errorf("result should have %d requests, but had %d", itemSize, len(result.Items))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}

}
