package view

import (
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/frontend/view/model"
	common "github.com/zhangzw001/crawler/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := template.Must(template.ParseFiles("template.html"))
	out , _ := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 12
	item := engine.Item{
		Url:  "http://www.youyuan.com/898155650-profile/",
		Type: "youyuan",
		Id:   "898155650",
		Payload: common.Profile{
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
			Sex:         "看情况",
		},
	}
	for i :=0 ; i < 12 ; i++ {
		page.Items = append(page.Items,item)
	}
	template.Execute(out, page)
}

func TestS(t *testing.T)  {
	view := CreateSearchResultView("template.html")

	//template:= template.Must(
	//template.ParseFiles("template.html"))

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}

	page.Hits = 123

	item := engine.Item{
		Url:  "http://www.youyuan.com/898155650-profile/",
		Type: "youyuan",
		Id:   "898155650",
		Payload: common.Profile{
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
			Sex:         "看情况",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items,item)
	}

	//err = template.Execute(out, page)
	err = view.Render(out,page)
	if err != nil{
		panic(err)
	}
}
