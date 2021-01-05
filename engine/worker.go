package engine

import (
	"github.com/zhangzw001/crawler/fetcher"
	"github.com/zhangzw001/crawler/public"
	"log"
)

func worker(req Request) (ParseResult,error ){
	// 获取url的内容
	result := ParseResult{}
	log.Printf("Fetching url : %v \n",req.Url)
	body , err := fetcher.Fetch(req.Url)
	if err != nil {
		return result,nil
	}
	// 对fetch的内容调用函数
	result = req.ParserFunc(body)
	return result,nil
}


func isDuplicated(url string) bool  {
	if public.Duplicated[url] {
		return true
	}
	public.Duplicated[url] = true
	return false
}
