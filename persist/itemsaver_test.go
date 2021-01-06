package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/model"
	"github.com/zhangzw001/crawler/public"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"testing"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url:  "http://www.youyuan.com/898155650-profile/",
		Type: "youyuan",
		Id:   "898155650",
		Payload: model.Profile{
			Name:        "风无痕",
			Age:         44,
			Height:      150,
			Weight:      190,
			Income:      "小于2000元",
			Marriage:    "离异",
			Education:   "硕士及以上",
			Occupation:  "其他",
			Hokou:       "四川",
			House:       "已购房",
			Car:         "",
			WorkAddress: "",
			Sex:         "看情况",
		},
	}
	err := Save(expected)
	if err != nil {
		log.Println(err)
	}
	// 在从es中查询
	client, err := elastic.NewClient(elastic.SetURL(public.EsUrl),elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	resp , err := client.Get().Index(public.EsIndex).Type(public.EsType).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var item engine.Item
	json.Unmarshal(*resp.Source,&item)

	// 这里如果不json解析, Payload 是个map结构
	item.Payload,err = model.FromJsonToProfile(item.Payload)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(item)
	//fmt.Println(expected)

	if item != expected{
		fmt.Println("error")
	}
}
