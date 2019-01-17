package persist

import (
	"testing"
	"crawl_zhenai2/model"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
	"crawl_zhenai2/engine"
)

func TestSave(t *testing.T) {
	// "basicInfo": ["未婚", "26岁", "魔羯座(12.22-01.19)", "165cm", "50kg", "工作地:苏州相城区", "月收入:5-8千", "职业技术教师", "高中及以下"],

	item := engine.Item{
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
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "datint_test"
	err = save(client,item,index)
	// 使用docker go client

	//从ElasticSearch中获取，根据id

	resp, err := client.Get().
		Index(index).
		Type(item.Type).
		Id(item.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source) //打印

	//反序列化
	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)

	if err != nil {
		panic(err)
	}

	actualProfile,_ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	//断言
	if actual != item {
		t.Errorf("got %v; expected %v", actual, item)
	}
}
