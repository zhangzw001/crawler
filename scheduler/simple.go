package scheduler

import "github.com/zhangzw001/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(req chan engine.Request) {
	s.WorkerChan = req
}

func (s *SimpleScheduler) Submit(req engine.Request) {
	go func(){
		s.WorkerChan <- req
	}()
}

