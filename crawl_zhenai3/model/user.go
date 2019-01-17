package model

import "encoding/json"

//珍爱网用户对象模型
type Profile struct {
	Name       string //姓名
	Marriage   string //婚况
	Age        string //年龄
	Gender     string //性别
	Height     string //身高
	Weight     string //体重
	Income     string //收入
	Education  string //教育
	Occupation string //职业
	Hukou      string //籍贯户口
	Xingzuo    string //星座
	House      string //房子
	Car        string //车
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err

}
