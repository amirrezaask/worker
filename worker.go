package worker

type Worker interface {
	RecentJobs(n int) map[string]JobStatus
	Failed() map[string]JobStatus
	QueueLen() int
	Start()
}

type simpleWorker struct {
	numOfGoroutines int
	queue           Queue
	errHandler      func(l Logger, err error)
	logger          Logger
	doneJobs        map[string]JobStatus
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
					continue
				}
				err = j.Run()
				if err != nil {
					w.errHandler(w.logger, err)
					continue
				}
				w.doneJobs[j.ID()] = j.Status()
			}
		}(w.queue)
	}
}

func (w *simpleWorker) RecentJobs(n int) map[string]JobStatus {
	out := make(map[string]JobStatus)
	var i int
	for id, status := range w.doneJobs {
		if i > n {
			break
		}
		out[id] = status
		i++
	}
	return out
}

func (w *simpleWorker) QueueLen() int {
	return w.queue.Len()
}

func (w *simpleWorker) Failed() map[string]JobStatus {
	out := make(map[string]JobStatus)
	for id, status := range w.doneJobs {
		if status == JobFinishedNoErrs || status == JobFinishedWithErrs {
			out[id] = status
		}
	}
	return out
}
