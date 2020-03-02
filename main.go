package main

import (
	"webcrawler/engine"
	"webcrawler/persist"
	"webcrawler/scheduler"
	"webcrawler/zhenai/parser"
)

func main() {
	itemSaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemSaver,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
