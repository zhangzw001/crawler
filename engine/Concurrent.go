package engine

import (
	"log"
)

//
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// createWorker 函数传入 Scheduler 比较重
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	//
	for _, req := range seeds {
		go e.Scheduler.Submit(req)
	}
	out := make(chan ParseResult)
	//
	go e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		go createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item : %v \n", item)
		}
		for _, req := range result.Requests {
			go e.Scheduler.Submit(req)
		}
	}

}

func createWorker(in chan Request, result chan ParseResult, ready ReadyNotifier) {
	for {
		// tel scheduler i'm ready
		ready.WorkerReady(in)
		request := <-in
		parseResult, err := worker(request)
		if err != nil {
			continue
		}
		result <- parseResult
	}
}
