package main

import (

	"regexp"
	"fmt"
	"io/ioutil"
)

func main() {

	//url := "http://www.zhenai.com/zhenghun"
	//engine.Run(engine.Request{
	//	Url:url,
	//	ParserFunc:parser.ParseCityList,
	//})



/*
	//url:="http://album.zhenai.com/u/1214814888"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code:", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
*/

	all, err := ioutil.ReadFile("parser/citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	PrintCityList(all)

}


//打印城市信息
func PrintCityList(contents []byte) {

	re := regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)
	//re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	all := re.FindSubmatch(contents)
	fmt.Println(all)
	fmt.Printf("%s",all[1])

}




