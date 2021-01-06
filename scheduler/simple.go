package scheduler

import "github.com/zhangzw001/crawler/engine"

type Simple struct {
	// 简单版本只需要一个chan 即可
	workerChan chan engine.Request
}

func CreateSimple() *Simple{
	var s Simple
	s.workerChan = make(chan engine.Request)
	return &s
}

func (s *Simple) Submit(req engine.Request) {
	s.workerChan <- req
}

func (s *Simple) GetWorkerChan() chan engine.Request {
	return s.workerChan
}

