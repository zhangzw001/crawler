package scheduler

import "github.com/zhangzw001/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request	//用户输入chan
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//
	//s.WorkerChan <- r
	go func() {
		s.WorkerChan <- r
	}()
}

