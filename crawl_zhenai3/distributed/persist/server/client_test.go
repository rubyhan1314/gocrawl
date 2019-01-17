package main

import (
	"testing"
	"crawl_zhenai3/distributed/rpcsupport"
	"crawl_zhenai3/engine"
	"crawl_zhenai3/model"
	"time"
	"crawl_zhenai3/distributed/config"
)

func TestItemSaver(t *testing.T) {

	const host = ":7788"

	//1. 启动ItemSave的rpc服务
	go serveRpc(host, "test1")
	time.Sleep(time.Second) //先暂停一下，让rpc服务，起来

	//2. 启动ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	//3. 调用save()
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1214814888",
		Type: "zhenai",
		Id:   "1214814888",
		Payload: model.Profile{
			Name:       "林YY",
			Marriage:   "未婚",
			Age:        "26岁",
			Xingzuo:    "魔羯座(12.22-01.19)",
			Height:     "165cm",
			Weight:     "50kg",
			Income:     "月收入:5-8千",
			Occupation: "职业技术教师",
			Education:  "高中及以下",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result :%s ;  err : %v", result, err)
	}

}
