package scheduler

import (
	"webcrawler/concurrent/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker channel
	go func() {
		s.workerChan <- r
	}()
}
