package model

import "encoding/json"

type Profile struct {
	Name        string //名字,昵称
	Gender      string //性别
	Age         int    //年龄
	Height      int    //身高
	Weight      int    //体重
	Income      string //收入
	Marriage    string //婚姻状况
	Education   string //教育
	Occupation  string //职业
	Hokou       string //户口,籍贯
	House       string //房子
	Car         string //车
	WorkAddress string //工作地点
	Sex         string //能否接受婚前性行为
}


func FromJsonToProfile(i interface{}) (Profile, error) {
	var p Profile
	b, err := json.Marshal(i)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(b,&p)
	return p, nil

}
