package parser

import (
	"crawl_zhenai3/engine"
	"regexp"
	"log"
	"github.com/bitly/go-simplejson"
	"crawl_zhenai3/model"
	"fmt"
	"strings"
	"strconv"
	"crawl_zhenai3/distributed/config"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		//fmt.Printf("json : %s\n",json)
		profile, id := parseJson(json)
		profile.Name = name
		//fmt.Println(profile)
		//result.Items = append(result.Items, profile)
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Type:    "zhenai",
			Id:      id,
			Payload: profile,
		})
	}

	return result

}

//解析json数据
func parseJson(json []byte) (model.Profile, string) {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()
	//infos是一个切片，里面的类型是interface{}

	//fmt.Printf("infos:%v,  %T\n", infos, infos) //infos:[离异 47岁 射手座(11.22-12.21) 157cm 55kg 工作地:阿坝汶川 月收入:3-5千 教育/科研 大学本科],  []interface {}

	var profile model.Profile
	//所以我们遍历这个切片，里面使用断言来判断类型

	len := len(infos)
	for k, v := range infos {
		//fmt.Printf("k:%v,%T\n", k, k)
		//fmt.Printf("v:%v,%T\n", v, v)

		/*
		 "basicInfo":[
            "未婚",
            "25岁",
            "魔羯座(12.22-01.19)",
            "152cm",
            "42kg",
            "工作地:阿坝茂县",
            "月收入:3-5千",
            "医生",
            "大专"
        ],
		 */
		if e, ok := v.(string); ok {

			if strings.Contains(e, "未婚") || strings.Contains(e, "离异") || strings.Contains(e, "丧偶") {
				profile.Marriage = e
			} else if strings.Contains(e, "岁") {
				profile.Age = e
			} else if strings.Contains(e, "座") {
				profile.Xingzuo = e
			} else if strings.Contains(e, "cm") {
				profile.Height = e
			} else if strings.Contains(e, "kg") {
				profile.Weight = e
			} else if strings.Contains(e, "月收入") {
				profile.Income = e
			} else{
				//剩下的两个内部不太好写，我们可以按照下标来解析
				switch k {
				case len-2:
					profile.Occupation = e
				case len-1:
					profile.Education = e
				}
			}


		}

	}

	infos2, err := res.Get("objectInfo").Get("detailInfo").Array()

	for _, v := range infos2 {
		/*
		"detailInfo": ["汉族", "籍贯:江苏宿迁", "体型:富线条美", "不吸烟", "不喝酒", "租房", "未买车", "没有小孩", "是否想要孩子:想要孩子", "何时结婚:认同闪婚"],
       因为每个 每个用户的detailInfo数据不同，我们可以通过提取关键字来判断
*/
		if e, ok := v.(string); ok {
			//fmt.Println(k, "--->", e)
			if strings.Contains(e, "族") {
				profile.Hukou = e
			} else if strings.Contains(e, "房") {
				profile.House = e
			} else if strings.Contains(e, "车") {
				profile.Car = e
			}
		}
	}

	//性别：

	gender, err := res.Get("objectInfo").Get("genderString").String()
	fmt.Println("gender:", gender)
	profile.Gender = gender

	//id
	id, err := res.Get("objectInfo").Get("memberID").Int()

	fmt.Printf("%+v\n", profile)
	return profile, strconv.Itoa(id)
}




type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents [] byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	//return "ProfileParser", p.userName
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
