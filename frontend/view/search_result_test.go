package view

import (
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/frontend/view/model"
	common "github.com/zhangzw001/crawler/model"
	"github.com/zhangzw001/crawler/public"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	//template:= template.Must(
	//template.ParseFiles("view/template.html"))

	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")


	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://www.youyuan.com/898155650-profile/",
		Type: public.EsType,
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
			Sex:         "看情况1",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items,item)
	}

	//err = template.Execute(out, page)
	// Render 就会执行 template.Execute
	err = view.Render(out,page)
	if err != nil {
		panic(err)
	}

}
