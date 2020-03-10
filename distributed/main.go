package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"net/rpc"
	"strings"
	concurrentConfig "webcrawler/concurrent/config"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/scheduler"
	"webcrawler/concurrent/uestcfaculty/parser"
	saver "webcrawler/distributed/persist/client"
	"webcrawler/distributed/rpcsupport"
	worker "webcrawler/distributed/worker/client"
)

var (
	itemSaverHost = flag.String("saver-host", "", "item saver host")
	workerHosts = flag.String(
		"worker-hosts", "", "comma separated worker hosts")
)

func main() {
	flag.Parse()
	itemSaver, err := saver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	clientPool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(clientPool)

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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s successfully", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
