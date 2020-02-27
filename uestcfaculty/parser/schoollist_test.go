package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParseSchoolList(t *testing.T) {
	resp, err := http.Get("http://faculty.uestc.edu.cn/xylb.jsp?urltype=tree.TreeTempUrl&wbtreeid=1021")
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("%v", err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%v", err)
	}
	//fmt.Printf("%s\n%d", string(contents), resp.StatusCode)
	s := ParseSchoolList(contents)
	fmt.Println(s)
}
