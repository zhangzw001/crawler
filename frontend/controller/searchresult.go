package controller

import (
	"context"
	"github.com/zhangzw001/crawler/engine"
	view2 "github.com/zhangzw001/crawler/frontend/view"
	"github.com/zhangzw001/crawler/frontend/view/model"
	"github.com/zhangzw001/crawler/public"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view2.SearchResultView
	client *elastic.Client
}

//初始化
func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetURL(public.EsUrl),elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view2.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=已购房&from=20
func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from,err := strconv.Atoi(req.FormValue("from"))
	//填错问题初始化0
	if err != nil {
		from = 0
	}

	//fmt.Fprintf(w,"q := %s, from := %d ",q,from)
	var page model.SearchResult
	page, err  = s.getSearchResult(q, from )
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
	// 这里Render就会更加查询得到的数据 生成template文件
	err = s.view.Render(w,page)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
}


func (s SearchResultHandler) getSearchResult(q string ,from int) (model.SearchResult,error ) {
	var result  model.SearchResult

	resp , err := s.client.Search(public.EsIndex).
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, nil
	}

	result.Hits = resp.TotalHits()
	// 方法一, 直接修改result.Items的类型 为[]interface{}
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	// 方法二, 如果不修改result.Items的类型
	//for _, v := range resp.Each(reflect.TypeOf(engine.Item{})) {
	//	item := v.(engine.Item)
	//	result.Items = item
	//}

	result.Start = from
	result.Query = q
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result,nil
}


func rewriteQueryString(q string) string {
	//log.Println(q)
	//q = strings.ReplaceAll(q," "," AND ")
	//log.Println(q)
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
