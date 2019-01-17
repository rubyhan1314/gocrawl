package main

import (
	"crawl_zhenai3/scheduler"
	"crawl_zhenai3/engine"
	"crawl_zhenai3/parser"
	"crawl_zhenai3/distributed/persist/client"
	"fmt"
	"crawl_zhenai3/distributed/config"
	client2 "crawl_zhenai3/distributed/worker/client"
	"net/rpc"
	"crawl_zhenai3/distributed/rpcsupport"
	"log"
	"flag"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHost    = flag.String("worker_hosts", "", "worker hosts (comma separated)") //逗号分隔
)

func main() {


	url := "http://www.zhenai.com/zhenghun"

	//itemChan, err := persist.ItemSaver("datint_profile")
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	flag.Parse()

	pool := createClientPool(strings.Split(*workerHost, ","))

	processor := client2.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: url,
		//ParserFunc: parser.ParseCityList,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients [] *rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s : %v", h, err)
		}
	}

	//分发
	out := make(chan *rpc.Client)
	go func() {

		for {
			for _, client := range clients {
				out <- client
			}
		}

	}()

	return out
}
