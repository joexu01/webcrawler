package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParseSchoolList(t *testing.T) {
	resp, err := http.Get("http://faculty.uestc.edu.cn/xylb.jsp?urltype=tree.TreeTempUrl&wbtreeid=1021")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	//fmt.Printf("%s\n%d", string(contents), resp.StatusCode)
	s := ParseSchoolList(contents)
	fmt.Printf("%+v", s)
}
