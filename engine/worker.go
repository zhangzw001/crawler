package engine

import (
	"github.com/zhangzw001/crawler/fetcher"
	"github.com/zhangzw001/crawler/public"
)

//将fetcher和parse合起来 成worker
//fetcher的输出刚好是parse函数的输入
func worker(req Request) (ParseResult,error ){
	// 获取url的内容
	body , err := fetcher.Fetch(req.Url)
	if err != nil {
		return ParseResult{},nil
	}
	// 对fetch的内容调用函数
	return req.ParserFunc(body),nil
}

func isDuplicated(url string) bool  {
	if public.Duplicated[url] {
		return true
	}
	public.Duplicated[url] = true
	return false
}
