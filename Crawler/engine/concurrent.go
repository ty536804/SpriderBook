package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item //数据存储
}

type Scheduler interface {
	ReadNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		CreateWork(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, m := range seeds {
		e.Scheduler.Submit(m)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
			//log.Printf("Got item: %v", item)
		}
		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult, ready ReadNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
