package parser

import (
	"crawl_zhenai2/engine"
	"regexp"
)

var (
	cityRe    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

//解析信息
func ParseCity(contents []byte) engine.ParseResult {

	all := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		//fmt.Println("用户url：", string(c[1]))
		//result.Items = append(result.Items, "User:"+string(c[2])) //用户名字

		url:=string(c[1])
		name := string(c[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				//return ParseProfile(c, name)
				return ParseProfile(c,url,name)
			},
		})
	}
	//爬取下一页
	all2 := cityUrlRe.FindAllSubmatch(contents, -1)
	for _,c :=range all2{
		result.Requests = append(result.Requests,engine.Request{
			Url:string(c[1]),
			ParserFunc:ParseCity,
		})

	}


	return result
}

