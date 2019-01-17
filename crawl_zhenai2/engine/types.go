package engine

import "crawl_zhenai2/model"

//解析后返回的结果
type ParseResult struct {
	Requests []Request
	//Items    []interface{}
	Items    []Item
}

type Request struct {
	Url        string                   //解析出来的URL
	ParserFunc func([]byte) ParseResult //处理这个URL所需要的函数
}


type Item struct {
	Url string //URL
	Type string //存储到ElasticSearch时的type
	Id  string //用户Id
	//Payload interface{}
	Payload model.Profile
}
func NilParser([] byte) ParseResult {
	return ParseResult{}
}
