package scheduler

import (
	"github.com/zhangzw001/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func CreateSimple() *SimpleScheduler {
	return &SimpleScheduler{
		workerChan: make(chan engine.Request),
	}
}
func (s *SimpleScheduler) Submit(req engine.Request) {
		s.workerChan <- req
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}


func (s *SimpleScheduler) Run() {
	// 在这里make会有冲突, 因为concurrent里面我是go Run() 执行
	//s.workerChan = make(chan engine.Request)
}
