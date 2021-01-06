package engine

import "fmt"

// 请求需要记录url
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// 返回结果是Request列表 和 item
type ParseResult struct {
	Requests []Request
	Items    []Item
}

func NilParser([]byte) ParseResult {
	fmt.Println("this is NilParser")
	return ParseResult{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}


