package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"webcrawler/distributed/config"
	"webcrawler/distributed/persist"
	"webcrawler/distributed/rpcsupport"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
