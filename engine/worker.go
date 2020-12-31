package engine

import (
	"github.com/zhangzw001/crawler/fetcher"
)

func worker(req Request) (ParseResult,error ){
	// 获取url的内容
	result := ParseResult{}
	body , err := fetcher.Fetch(req.Url)
	if err != nil {
		return result,nil
	}
	// 对fetch的内容调用函数
	result = req.ParserFunc(body)
	return result,nil
}
