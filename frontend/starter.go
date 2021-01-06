package main

import (
	"github.com/zhangzw001/crawler/frontend/controller"
	"net/http"
)


func main() {
	http.Handle("/",http.FileServer(http.Dir("frontend/view")))
	//http.Handle("/search", controller.SearchResultHandler{})
	http.Handle("/search", controller.CreateSearchResultHandler("frontend/view/template.html"))

	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
}
