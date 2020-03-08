package main

import (
	"testing"
	"time"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/model"
	"webcrawler/distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// Start ItemSaver Server
	go serveRpc(host, "saver_test")
	time.Sleep(time.Second)

	// Start ItemSaver Client
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call Save Service
	item := engine.Item{
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
	var result string
	err = client.Call("ItemSaveService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s, error: %s", result, err)
	}
}
