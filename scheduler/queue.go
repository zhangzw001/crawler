package scheduler

import (
	"github.com/zhangzw001/crawler/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func CreateQueue() *QueueScheduler {
	var q QueueScheduler
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	return &q
}
func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {
	//s.requestChan = make(chan engine.Request)
	//s.workerChan = make(chan chan engine.Request)

	var requestQ []engine.Request
	var workerQ []chan engine.Request

	for {
		var activeRequest engine.Request
		var activeWorker chan engine.Request
		if len(requestQ) > 0 && len(workerQ) > 0 {
			activeRequest = requestQ[0]
			activeWorker = workerQ[0]
		}
		select {
		case r := <-s.requestChan:
			requestQ = append(requestQ, r)
		case w := <-s.workerChan:
			workerQ = append(workerQ, w)
		case activeWorker <- activeRequest:
			requestQ = requestQ[1:]
			workerQ = workerQ[1:]
		}
	}
}
