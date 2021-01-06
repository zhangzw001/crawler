package engine

import (
	"log"
)

type Scheduler interface {
	Submit(Request)
	GetWorkerChan() chan Request
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//1. 爬取种子
	for _, req := range seeds {
		// ? Submit 应该做什么呢?
		// 之前单任务的时候是添加到一个list
		// 这里显然直接写入到 chan 即可 , chan就像一个队列, 顺序执行
		go e.Scheduler.Submit(req)
	}

	//2. 创建worker
	in := e.Scheduler.GetWorkerChan()
	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		go createWorker(in, out)
	}

	//3. 对结果继续爬取
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item : %v\n", item)
		}

		for _, req := range result.Requests {
			go e.Scheduler.Submit(req)
		}
	}
}

// 2. 创建worker的函数
func createWorker(in chan Request, out chan ParseResult) {
	for {
		req := <-in
		log.Printf("Working url : %v \n",req.Url)
		result, err := worker(req)
		if err != nil {
			continue
		}
		out <- result
	}
}
