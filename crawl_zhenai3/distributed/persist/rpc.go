package persist

import (
	"crawl_zhenai3/engine"
	"github.com/olivere/elastic"
	"crawl_zhenai3/persist"
	"log"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	log.Printf("Item %v saved..", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v : %v", item, err)
	}
	return err
}
