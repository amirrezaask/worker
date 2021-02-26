package worker

import "fmt"

type Worker interface {
	RecentJobs(n int) map[string]JobStatus
	Failed() map[string]JobStatus
	QueueLen() int
	Start()
	RegisterJob(j Job)
	RunJob(j Job)
	RunJobByID(jobID string)
}

type simpleWorker struct {
	numOfGoroutines int
	queue           Queue
	errHandler      func(l Logger, err error)
	logger          Logger
	doneJobs        map[string]JobStatus
	jobMap          map[string]Job
}

func NewSimpleWorker(numOfGoroutines int, queue Queue) Worker {
	return &simpleWorker{numOfGoroutines: numOfGoroutines, queue: queue}
}

func (w *simpleWorker) Start() {
	for i := 0; i < w.numOfGoroutines; i++ {
		go func(q Queue) {
			for {
				jobID, err := q.Next()
				if err != nil {
					w.errHandler(w.logger, err)
					continue
				}
				j, exists := w.jobMap[jobID]
				if !exists {
					w.errHandler(w.logger, fmt.Errorf("no job with id %s is registered", jobID))
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

func (w *simpleWorker) RegisterJob(j Job) {
	w.jobMap[j.ID()] = j
}

func (w *simpleWorker) jobIsRegistered(j Job) bool {
	_, exists := w.jobMap[j.ID()]
	return exists
}

func (w *simpleWorker) RunJob(j Job) {
	if !w.jobIsRegistered(j) {
		w.RegisterJob(j)
	}
	w.queue.Add(j.ID())
}

func (w *simpleWorker) RunJobByID(jobID string) {
	w.queue.Add(jobID)
}
