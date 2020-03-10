package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"webcrawler/distributed/config"
	"webcrawler/distributed/persist"
	"webcrawler/distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("ItemSaver: must specify a port.\n")
		fmt.Println("Type --help for more information")
		return
	}
	log.Printf("ItemSaver service listening on %d", *port)
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	})
}
