package client

import (
	"crawl_zhenai3/engine"
	"log"
	"crawl_zhenai3/distributed/rpcsupport"
	"crawl_zhenai3/distributed/config"
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
			log.Printf("Item Saver: Got %d  item : %v\n", itemCount, item)
			itemCount++

			//调用Rpc 来保存item

			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)

			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}

		}
	}()
	return out, nil

}
