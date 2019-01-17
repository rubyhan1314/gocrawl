package main

import (
	"testing"
	"crawl_zhenai3/distributed/rpcsupport"
	"crawl_zhenai3/distributed/worker"
	"time"
	"crawl_zhenai3/distributed/config"
	"fmt"
)

func TestCrawlService(t *testing.T) {
	const host = ":5566"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1214814888",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "林YY",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
