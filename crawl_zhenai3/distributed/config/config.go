package config

const (
	// Service端口
	ItemSaverPort = 7788 //ItemSaver RPC远程服务的端口
	WorkerPort0   = 5566 //Worker

	//RPC远程服务要调用的方法名称
	ItemSaverRpc    = "ItemSaveService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	//Parser names
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	ElasticIndex = "dating_profile" //ElasticSearch中要存储的Index名称

)
