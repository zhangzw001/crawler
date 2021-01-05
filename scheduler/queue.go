package scheduler

import "github.com/zhangzw001/crawler/engine"

type QueueScheduler struct {

}

func (q QueueScheduler) Submit(engine.Request) {
	panic("implement me")
}

func (q QueueScheduler) ConfigMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}
