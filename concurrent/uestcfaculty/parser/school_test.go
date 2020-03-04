package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseSchool(t *testing.T) {
	//resp, err := http.Get("http://faculty.uestc.edu.cn/xylb.jsp?urltype=tsites.CollegeTeacherList&wbtreeid=1021&st=0&id=2031&lang=zh_CN#collegeteacher")
	//if err != nil {
	//	t.Errorf("%v", err)
	//}
	//defer func() {
	//	err := resp.Body.Close()
	//	if err != nil {
	//		t.Errorf("an error occured when closing resp.Body: %v", err)
	//	}
	//}()
	var urls = []string{
		"http://faculty.uestc.edu.cn/caihongbin/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/caishimin/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/chenaiguo/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/chenduanbing/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/chenfu/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/chenjuan1/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/chenting/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/daibo/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/daiyuanshun/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/dengjian/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/dingxuyang/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/dongle/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/dongqiang/zh_CN/index.htm",
		"http://faculty.uestc.edu.cn/xylb.jsp?totalpage=9&PAGENUM=2&urltype=tsites.CollegeTeacherList&wbtreeid=1021&st=0&id=2031&lang=zh_CN",
	}
	contents, err := ioutil.ReadFile("./test_files/school_test.txt")
	if err != nil {
		t.Errorf("%v", err)
	}
	result := ParseSchool(contents, "")
	for i, r := range result.Requests {
		if r.Url == urls[i] {
			fmt.Printf("got url: %s\n", r.Url)
		} else {
			t.Errorf("supposed: %s; but got: %s", urls[i], r.Url)
		}
	}

}
