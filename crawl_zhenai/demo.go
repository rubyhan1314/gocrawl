package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
)

var json_str string = `{"rc" : 0,
  "error" : "Success",
  "type" : "stats",
  "progress" : 100,
  "job_status" : "COMPLETED",
  "result" : {
    "total_hits" : 803254,
    "starttime" : 1528434707000,
    "endtime" : 1528434767000,
    "fields" : [ ],
    "timeline" : {
      "interval" : 1000,
      "start_ts" : 1528434707000,
      "end_ts" : 1528434767000,
      "rows" : [ {
        "start_ts" : 1528434707000,
        "end_ts" : 1528434708000,
        "number" : "x12887"
      }, {
        "start_ts" : 1528434720000,
        "end_ts" : 1528434721000,
        "number" : "x13028"
      }, {
        "start_ts" : 1528434721000,
        "end_ts" : 1528434722000,
        "number" : "x12975"
      }, {
        "start_ts" : 1528434722000,
        "end_ts" : 1528434723000,
        "number" : "x12879"
      }, {
        "start_ts" : 1528434723000,
        "end_ts" : 1528434724000,
        "number" : "x13989"
      } ],
      "total" : 803254
    },
      "total" : 8
  }
}`

func main() {

	res, err := simplejson.NewJson([]byte(json_str))

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	//获取json字符串中的 result 下的 timeline 下的 rows 数组
	rows, err := res.Get("result").Get("timeline").Get("rows").Array()


	fmt.Printf("rows:%T\n",rows)
	//遍历rows数组
	for _, row := range rows {
		fmt.Printf("rows:%T\n",row)
		//对每个row获取其类型，每个row相当于 C++/Golang 中的map、Python中的dict
		//每个row对应一个map，该map类型为map[string]interface{}，也即key为string类型，value是interface{}类型
		if each_map, ok := row.(map[string]interface{}); ok {

			//可以看到each_map["start_ts"]类型是json.Number
			//而json.Number是golang自带json库中decode.go文件中定义的: type Number string
			//因此json.Number实际上是个string类型
			fmt.Println(reflect.TypeOf(each_map["start_ts"]))

			if start_ts, ok := each_map["start_ts"].(json.Number); ok {
				start_ts_int, err := strconv.ParseInt(string(start_ts), 10, 0)
				if err == nil {
					fmt.Println(start_ts_int)
				}
			}

			if number, ok := each_map["number"].(string); ok {
				fmt.Println(number)
			}

		}
	}
}