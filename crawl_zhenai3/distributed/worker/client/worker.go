package client

import (
	"crawl_zhenai3/engine"
	"crawl_zhenai3/distributed/config"
	"crawl_zhenai3/distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor) {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		client := <-clientChan

		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
