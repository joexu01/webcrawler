package parser

import (
	"io/ioutil"
	"testing"
	"webcrawler/engine"
	"webcrawler/model"
)

func TestParseTeacherProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./test_files/teacher_test.txt")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	supposed := engine.Item{
		Url: "http://faculty.uestc.edu.cn/chenaiguo/zh_CN/index.htm",
		Id:  "chenaiguo",
		Payload: model.TeacherProfile{
			Name:          "陈爱国",
			Gender:        "男",
			GraduatedFrom: "北京邮电大学",
			Position:      "研究员",
			Education:     "博士研究生毕业",
			Degree:        "工学博士学位",
			IsIncumbent:   "在岗",
			School:        "计算机科学与工程学院（网络空间安全学院）",
			Discipline:    "",
			Location:      "清水河校区主楼A2-301B",
			Email:         "",
			PersonalUrl:   "http://faculty.uestc.edu.cn/chenaiguo/zh_CN/index.htm",
		},
	}

	result := ParseTeacherProfile(contents,
		"陈爱国", "研究员", "http://faculty.uestc.edu.cn/chenaiguo/zh_CN/index.htm")
	if supposed != result.Items[0] {
		t.Errorf("an error occued when verifying results: got %+v, supposed: %+v", result.Items[0], supposed)
	}
}
