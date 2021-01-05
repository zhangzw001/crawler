package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
}

// 并发版engine, 通过 Scheduler 接口实现
// main函数传入具体的scheduler的类型,实现 interface的功能
// 1. simple类型 SimpleScheduler
// 2. queue类型  QueueScheduler
func (e *ConcurrentEngine) Run(seeds ...Request) {
	//  开始建worker
	//  新建两个channel, 一个输入 一个输出
	in := make(chan Request)
	out := make(chan ParseResult)
	// Scheduler.WorkerChan = in
	// 因为 in 其实就是 Scheduler.WorkerChan, 所以 worker 会等待读 in chan的数据, in什么时候有数据呢? 答案是 Submit 的时候(s.WorkerChan <- r)
	// 相当于 scheduler 送给 worker
	// 因此 worker 等待 in 的输入, 数据来自 scheduler 的 Submit
	// Submit 的 输入 req 是来自engine的 request
	// engine 的 request 除了第一次是 seeds 传入, 之后都是来自 worker 返回的 out
	e.Scheduler.ConfigMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		// 架构第一步 这里相当于 request 送给 worker(fetcher+parse) 返回给 engine 的 chan
		createWorker(in, out)
	}

	// 在scheduler中添加
	for _, req := range seeds {
		// 就是把请求req 都送到 Scheduler 的 WorkerChan 里面
		// 这里用go func() 快速结束
		// 架构第二步 这里相当于  request 送给 scheduler chan
		e.Scheduler.Submit(req)
	}
	// 最后对worker的结果进行处理
	itemCount := 0
	for {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item #%d:%v\n", itemCount, item)
			itemCount++
		}

		for _, req := range result.Requests {
			// 这里用go func() 快速结束
			e.Scheduler.Submit(req)
		}
	}
}

// createWorker 主要是 调用worker进程 fetcher和parse, 通过chan来做到并发返回
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			// 读取输入的channel数据
			req := <-in
			log.Printf("Fetching url %s\n", req.Url)

			// 读取到的数据 传入worker
			result, err := worker(req)

			if err != nil {
				continue
			}
			// worker返回的结果送到 out channel
			//go func() {
			//	out <- result
			//}()
			out <- result
		}
	}()
}
