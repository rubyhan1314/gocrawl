package parser

import (
	"regexp"
	"crawl_zhenai3/engine"
	"github.com/bitly/go-simplejson"
	"log"
	"crawl_zhenai3/distributed/config"
)

//const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
const cityListRe = `<script>window.__INITIAL_STATE__=(.+);\(function`

//解析城市列表信息
func ParseCityList(contents []byte,_ string) engine.ParseResult {
//func ParseCityList(contents []byte) engine.ParseResult {


	//fmt.Println(string(contents))

	re := regexp.MustCompile(cityListRe)
	//all := re.FindAllSubmatch(contents, -1)
	json := re.FindSubmatch(contents)
	//fmt.Printf("---%s\n",all[1])

	all := parseJsonCityList(json[1])
	result := engine.ParseResult{}
	//i := 0

	for _, c := range all {
		//result.Items = append(result.Items, string(c[2])) //城市名字
		//fmt.Println(string(c[0]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			//ParserFunc: engine.NilParser,
			//ParserFunc: ParseCity,
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}

//解析json数据
func parseJsonCityList(json []byte) ([][]string) {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, _ := res.Get("cityListData").Get("cityData").Array()
	//infos是一个切片，里面的类型是interface{}

	var dataList [][]string
	//所以我们遍历这个切片，里面使用断言来判断类型
	for _, v := range infos {
		if each_map, ok := v.(map[string]interface{}); ok {
			//fmt.Println(each_map)
			map2 := each_map["cityList"]
			for _, v2 := range map2.([]interface{}) {
				if data, ok := v2.(map[string]interface{}); ok {
					var datas [] string
					cityName := data["linkContent"].(string)
					cityUrl := data["linkURL"].(string)
					datas = append(datas, cityName, cityUrl)
					dataList = append(dataList, datas)
				}
			}
		}
	}
	return dataList

}
