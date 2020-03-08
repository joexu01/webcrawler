package main

import (
	"fmt"
	concurrentConfig "webcrawler/concurrent/config"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/scheduler"
	"webcrawler/concurrent/uestcfaculty/parser"
	"webcrawler/distributed/config"
	saver "webcrawler/distributed/persist/client"
	worker "webcrawler/distributed/worker/client"
)

func main() {
	itemSaver, err := saver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		RequestProcessor: processor,
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemSaver,
	}

	e.Run(engine.Request{
		Url:    "http://faculty.uestc.edu.cn/xylb.jsp?urltype=tree.TreeTempUrl&wbtreeid=1021",
		Parser: engine.NewFuncParser(parser.ParseSchoolList, concurrentConfig.ParseSchoolList),
	})
}
