package main

import (
	"flag"
	"fmt"
	"log"
	"webcrawler/distributed/rpcsupport"
	"webcrawler/distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("Worker: must specify a port.\n")
		fmt.Println("Type --help for more information")
		return
	}
	log.Fatal(
		rpcsupport.ServeRpc(
			fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
