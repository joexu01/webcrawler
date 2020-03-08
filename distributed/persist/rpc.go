package persist

import (
	"github.com/olivere/elastic/v7"
	"log"
	"webcrawler/concurrent/engine"
	"webcrawler/concurrent/persist"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
		log.Printf("item %v saved.", item)
	} else {
		log.Printf("error saving item %v: %v", item, err)
	}
	return err
}
