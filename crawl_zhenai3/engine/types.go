package engine

import (
	"crawl_zhenai3/model"
	"crawl_zhenai3/distributed/config"
)

//解析后返回的结果
type ParseResult struct {
	Requests []Request
	//Items    []interface{}
	Items []Item
}

type Request struct {
	Url string //解析出来的URL
	//ParserFunc func([]byte) ParseResult //处理这个URL所需要的函数
	Parser Parser
}

type Item struct {
	Url  string //URL
	Type string //存储到ElasticSearch时的type
	Id   string //用户Id
	//Payload interface{}
	Payload model.Profile
}

//func NilParser([] byte) ParseResult {
//	return ParseResult{}
//}

type Parser interface {
	Parse(contents [] byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type NilParser struct {
}

func (NilParser) Parse(_ [] byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (_ string, _ interface{}) {
	return config.NilParser, nil
}


type ParserFunc func(contents []byte, url string) ParseResult

type FuncParser struct {
	parserFunc ParserFunc
	name       string
}

func (f *FuncParser) Parse(contents [] byte, url string) ParseResult {
	return f.parserFunc(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parserFunc: p,
		name:       name,
	}
}

