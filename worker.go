package worker

type Worker interface {
	// Jobs() map[string]JobStatus
	Start()
}

type simpleWorker struct {
	numOfGoroutines int
	queue           Queue
	errHandler      func(l Logger, err error)
	logger          Logger
}

func NewSimpleWorker(numOfGoroutines int, queue Queue) Worker {
	return &simpleWorker{numOfGoroutines: numOfGoroutines, queue: queue}
}

func (w *simpleWorker) Start() {
	for i := 0; i < w.numOfGoroutines; i++ {
		go func(q Queue) {
			for {
				j, err := q.Next()
				if err != nil {
					w.errHandler(w.logger, err)
				}
				err = j.Run()
				if err != nil {
					w.errHandler(w.logger, err)
				}
			}
		}(w.queue)
	}
}

// func (w *simpleWorker) Jobs() map[string]JobStatus {
// }
