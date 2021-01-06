package engine

import (
	"fmt"
	"log"
)

// 单机版的简单engine
type SingleEngine struct {}
// SingleEngine 主要是修改一下功能:
// 1. 提取fetch 和抓取函数 worker
func (s SingleEngine) Run(seeds ...Request) {
	// 申明一个需要爬取的请求队列
	var requests []Request
	// 将传入的 根Request请求添加到队列
	for _, req := range seeds {
		// ?? 之前计划在爬取之前, 验证是否爬取过, 但这样会导致列表内容还是重复
		// 因此这里, 在加入到列表之前验证是否存在列表中
		if isDuplicated(req.Url) {
			continue
		}
		requests = append(requests,req)
	}

	// 只要队列长度大于0 , 就继续爬取
	for len(requests) >  0 {
		// 怎么爬取呢?
		// 一个个取, 取第一个, 然后删除第一个
		req := requests[0]
		requests = requests[1:]
		//log.Printf("Fetching url %s\n",  req.Url)
		parseResult , err := worker(req)
		if err !=nil {
			log.Printf("Fetcher.Fetch err : %v", err)
		}
		for _ , item := range parseResult.Items {
			fmt.Printf("Got Item : %v\n",item)
		}
		// 再将返回的内容中的request 加到队列中
		for _, req := range parseResult.Requests {
			// 同样,这里在加入到列表之前验证是否存在列表中
			if isDuplicated(req.Url) {
				continue
			}
			requests = append(requests,req)
		}

	}
}
