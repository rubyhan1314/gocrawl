package parser

import (
	"testing"
	"io/ioutil"
	"crawl_zhenai3/model"
	"crawl_zhenai3/engine"
)

func TestParseProfile(t *testing.T) {
	contents,err :=ioutil.ReadFile("user_data.html")
	if err != nil{
		panic(err)
	}

	result := ParseProfile(contents,"http://album.zhenai.com/u/1214814888","林YY")

	if len(result.Items) != 1{
		t.Errorf("Items should contain 1 element; but was %v",result.Items)
	}

	actual:= result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1214814888",
		Type: "zhenai",
		Id:   "1214814888",
		Payload: model.Profile{
			Name:       "林YY",
			Marriage:   "未婚",
			Age:        "26岁",
			Xingzuo:    "魔羯座(12.22-01.19)",
			Height:     "165cm",
			Weight:     "50kg",
			Income:     "月收入:5-8千",
			Occupation: "职业技术教师",
			Education:  "高中及以下",
		},
	}


	if actual != expected{
		t.Errorf("expected %v , but was %v \n",expected,actual)
	}

}
