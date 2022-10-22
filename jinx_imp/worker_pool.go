package jinx_imp

import (
	"fmt"
	"jinx/jinx_int"
	"jinx/jinx_pkg/log"
)

type job struct {
	handler jinx_int.IMsgHandle
	req     jinx_int.IRequest
}

func (j *job) handle() {
	j.handler.PreHandle(j.req)
	j.handler.Handle(j.req)
	j.handler.PostHandle(j.req)
}

type workerPool struct {
	workerSize     int
	jobChannelSize int
	jobChannel     chan *job
	closeChan      chan struct{}
}

func (w *workerPool) Start() {
	for i := 0; i < w.workerSize; i++ {
		num := i
		w.startWorker(num)
	}
}

func (w *workerPool) startWorker(num int) {
	go func() {
		log.Info(fmt.Sprintf("worker:%v start...", num), ModuleNameWorker)
		for true {
			select {
			case <-w.closeChan:
				if len(w.jobChannel) > 0 {
					continue
				}
				log.Info(fmt.Sprintf("worker:%v closed...", num), ModuleNameWorker)
				return
			case j := <-w.jobChannel:
				j.handle()
			}
		}
	}()
}

func (w *workerPool) Submit(handle jinx_int.IMsgHandle, request jinx_int.IRequest) {
	j := &job{
		handler: handle,
		req:     request,
	}
	w.jobChannel <- j
}

func (w *workerPool) Stop() {
	close(w.closeChan)
}

func (w *workerPool) GetWorkerSize() int {
	return w.workerSize
}

func (w *workerPool) GetJobChannelSize() int {
	return w.jobChannelSize
}

func NewWorkerPool(workerSize, chanSize int) jinx_int.IWorkerPool {
	return &workerPool{
		workerSize:     workerSize,
		jobChannelSize: chanSize,
		jobChannel:     make(chan *job, chanSize),
		closeChan:      make(chan struct{}),
	}
}
