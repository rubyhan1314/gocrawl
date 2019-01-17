package parser

import (
	"testing"
	"crawl_zhenai3/fetcher"
)

func TestParseCityList(t *testing.T){
	contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	/*
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}
	*/

	parseResult := ParseCityList(contents)

	const resultSize = 470
	if len(parseResult.Requests) != resultSize{
		t.Errorf("result should have %d requests, but had %d",resultSize,len(parseResult.Requests))
	}

	//if len(parseResult.Items) != resultSize{
	//	t.Errorf("result should have %d Items, but had %d",resultSize,len(parseResult.Items))
	//}
}