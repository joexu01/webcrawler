package parser

import (
	"io/ioutil"
	"testing"
	"webcrawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		t.Errorf("an error occured when opening test file ./profile_test_data.html")
		return
	}

	result := ParseProfile(contents, "你值得拥有")
	if len(result.Items) != 1 {
		t.Errorf(
			"result.Items should contain 1 element, but had %d elements", len(result.Items))
	}
	profile := result.Items[0].(model.Profile)

	expectedProfile := model.Profile{
		Name:     "你值得拥有",
		Age:      24,
		Height:   168,
		Location: "阿坝九寨沟",
		Income:   "3千以下",
	}
	if profile != expectedProfile {
		t.Errorf("profiles not match")
	}
}
