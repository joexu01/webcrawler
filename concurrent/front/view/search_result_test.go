package view

import (
	"os"
	"testing"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/front/model"
	common "webcrawler/concurrent/model"
)

func TestSearchResultView_Render(t *testing.T) {
	v := CreateSearchResultView("template.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http://faculty.uestc.edu.cn/chenaiguo/zh_CN/index.htm",
		Id:  "chenaiguo",
		Payload: common.TeacherProfile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	out, err := os.Create("template.test.html")
	err = v.Render(out, page)
	if err != nil {
		panic(err)
	}
}
