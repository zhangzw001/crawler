package engine

import (
	"github.com/zhangzw001/crawler/scheduler"
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
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//
	for _, req  := range seeds {
		e.Scheduler.Submit(req)
	}

	//
	out := make(chan ParseResult)
	// 创建
	e.Scheduler = scheduler.CreateQueue()
	// run
	go e.Scheduler.Run()
	for i := 0 ; i < e.WorkerCount ; i ++ {
		go createWorker(out , e.Scheduler)
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



func createWorker(result chan ParseResult,s Scheduler) {
	in := make(chan Request)
	for {
		// tel scheduler i'm ready
		s.WorkerReady(in)
		request := <- in
		parseResult, err := worker(request)
		if err != nil {
			continue
		}
		result <- parseResult
	}
}
