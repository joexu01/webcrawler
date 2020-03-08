package main

import (
	"fmt"
	"testing"
	"time"
	"webcrawler/distributed/config"
	"webcrawler/distributed/rpcsupport"
	"webcrawler/distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url:    "http://faculty.uestc.edu.cn/chenaiguo/zh_CN/index.htm",
		Parser: worker.SerializedParser{
			Name: config.ParseTeacherProfile, Args: []string{"陈爱国", "研究员"},},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
