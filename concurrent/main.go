package main

import (
	"webcrawler/concurrent/config"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/persist"
	"webcrawler/concurrent/scheduler"
	"webcrawler/concurrent/uestcfaculty/parser"
)

func main() {
	itemSaver, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		RequestProcessor: engine.Worker,
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemSaver,
	}

	e.Run(engine.Request{
		Url:    "http://faculty.uestc.edu.cn/xylb.jsp?urltype=tree.TreeTempUrl&wbtreeid=1021",
		Parser: engine.NewFuncParser(parser.ParseSchoolList, config.ParseSchoolList),
	})
}
