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

	//
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
