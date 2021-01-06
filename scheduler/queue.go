package scheduler

import "github.com/zhangzw001/crawler/engine"

type Queue struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request // 因为会有很多个worker, 每一个worker的输入都是 chan Request
}
func CreateQueue() *Queue {
	var q Queue
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	return &q
}


func (q *Queue) CreateWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}


func (q *Queue) Submit(req engine.Request) {
	q.requestChan <- req
}

func (q *Queue) WorkerReady(in chan engine.Request){
	q.workerChan <- in
}

// Run的作用是协调不同的 workerChan 和 requestChan
// 队列的话 我们需要一个是 workerQueue 一个是 requestQueue
func (q *Queue) Run() {
	var workerQueue []chan engine.Request
	var requestQueue []engine.Request

	for {
		var activeWorker chan engine.Request
		var activeRequest engine.Request
		if len(workerQueue) > 0 && len(requestQueue) > 0 {
			// req队列不是空,并且worker队列也可以提供服务
			activeWorker = workerQueue[0]
			activeRequest = requestQueue[0]
			// 这里不能减少队列, 毕竟不一定能select到
		}
		select {
		// 如果requestChan 里面有值, 就加到队列
		case r := <-q.requestChan:
			requestQueue = append(requestQueue, r)
		case w := <-q.workerChan:
			workerQueue = append(workerQueue, w)
			// 只有request 真的被 worker 拿走之后, 才去掉队列
		case activeWorker <- activeRequest:
			workerQueue = workerQueue[1:]
			requestQueue = requestQueue[1:]
		}
	}
}
