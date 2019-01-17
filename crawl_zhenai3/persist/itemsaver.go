package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
	"crawl_zhenai3/engine"
	"errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	//关闭内网的sniff

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got %d  item : %v\n", itemCount, item)
			itemCount++

			err := Save(client, item,index) //保存item
			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}

		}
	}()
	return out, nil

}

//保存item
func Save(client *elastic.Client, item engine.Item, index string) error {
	//fmt.Printf("Save..%+v\n", item)

	if item.Type == "" {
		return errors.New("must supply Type ..")
	}

	indexService := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
		Index(index).
	//Type("zhenai").
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	//fmt.Printf("%+v",resp)//格式化输出结构体对象的时候包含了字段名称
	return nil
}
