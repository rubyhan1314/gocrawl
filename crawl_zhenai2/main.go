package main

import (

	"crawl_zhenai2/scheduler"
	"crawl_zhenai2/persist"
	"crawl_zhenai2/engine"
	"crawl_zhenai2/parser"
)

func main() {


	url := "http://www.zhenai.com/zhenghun"

	itemChan, err := persist.ItemSaver("datint_profile")
	if err != nil{
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:itemChan,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})


}