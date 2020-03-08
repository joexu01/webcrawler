package main

import (
	"fmt"
	"log"
	"webcrawler/distributed/config"
	"webcrawler/distributed/rpcsupport"
	"webcrawler/distributed/worker"
)

func main() {
	log.Fatal(
		rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
