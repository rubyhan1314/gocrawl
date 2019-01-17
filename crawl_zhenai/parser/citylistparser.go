package parser

import (
	"regexp"
	"crawl_zhenai/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

//解析城市信息
func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	i := 0
	for _, c := range all {
		result.Items = append(result.Items, string(c[2])) //城市名字

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			//ParserFunc: engine.NilParser,
			ParserFunc: ParseCity,
		})
		i++
		if i == 10{
			break
		}
	}

	return result
}




