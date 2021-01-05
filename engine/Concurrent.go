package engine

import (
	"log"
)

//
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}


type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request )
}

func (e ConcurrentEngine) Run(seeds ...Request) {

	//
	for _, req  := range seeds {
		e.Scheduler.Submit(req)
	}

	// Scheduler.WorkerChan = in
	// 因为 in 其实就是 Scheduler.WorkerChan, 所以 worker 会等待读 in chan的数据, in什么时候有数据呢? 答案是 Submit 的时候(s.WorkerChan <- r)
	// 相当于 scheduler 送给 worker
	// 因此 worker 等待 in 的输入, 数据来自 scheduler 的 Submit
	// Submit 的 输入 req 是来自engine的 request
	// engine 的 request 除了第一次是 seeds 传入, 之后都是来自 worker 返回的 out
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)
	for i := 0 ; i < e.WorkerCount ; i ++ {
		go createWorker(in, out )
	}

	for {
		result :=  <- out

		for _, item := range result.Items {
			log.Printf("Got item : %v \n",item)
		}

		for _, req := range result.Requests {
			e.Scheduler.Submit(req )
		}
	}

}



func createWorker(req chan Request, result chan ParseResult) {
	for {
		request := <-req
		parseResult, err := worker(request)
		if err != nil {
			continue
		}
		result <- parseResult
	}
}
