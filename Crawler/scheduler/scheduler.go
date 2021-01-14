package scheduler

import "Book/Crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(e engine.Request) {
	//panic("implement me")
	go func() {
		s.workerChan <- e
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	//panic("implement me")
	s.workerChan = r
}
