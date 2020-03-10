package client

import (
	"log"
	"webcrawler/concurrent/engine"
	"webcrawler/distributed/config"
	"webcrawler/distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver: got item #%d: %v\n", itemCount, item)
			itemCount++

			// call rpc to save item
			var result string
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v\n", item, err)
			}
		}
	}()
	return out, nil
}
