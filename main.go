package main

import (
	"github.com/zhangzw001/crawler/engine"
	"github.com/zhangzw001/crawler/scheduler"
	"github.com/zhangzw001/crawler/youyuan/parser"
)

const (
	urlYouYuan = "http://www.youyuan.com/city/"
	urlYouYuanCity = "http://www.youyuan.com/shanghai/"

)
func main() {
	// 第一步, 配置好 engine
	// 第二步, 配置好 ParserFunc, 每个request都有自己的 ParserFunc
	//  这里第一次配置一次 NilParser, 不做后续处理
	//  首页传入 "http://www.youyuan.com/city/" 城市页面, 会执行 CityListParser
	//  从city列表页面获取的 "http://www.youyuan.com/shanghai/" 上海页面, 会执行 CityParser
	//  最后从city页面获取 "http://www.youyuan.com/shanghai/xxx-profile" 某个用户的页面, 会执行 ProfileParser
	//engine.Run(engine.Request{
	//	Url:        urlYouYuan,
	//	ParserFunc: parser.CityListParser,
	//})


	e := engine.ConcurrentEngine{
		Scheduler: scheduler.CreateSimple(),
		WorkerCount: 3 ,
	}

	e.Run(engine.Request{
		Url:        urlYouYuan,
		ParserFunc: parser.CityListParser,
	})

	//e := engine.ConcurrentEngine{
	//	Scheduler: scheduler.CreateQueue(),
	//	WorkerCount: 3 ,
	//}
	//
	//e.Run(engine.Request{
	//	Url:        urlYouYuan,
	//	ParserFunc: parser.CityListParser,
	//})
}
