package main

import "github.com/amirrezaask/worker"

func main() {
	q := worker.NewChannelQueue(10)
	q.Add(worker.NewSimpleJob(func() error {
		return nil
	}))
	w := worker.NewSimpleWorker(2, q)
	w.Start()
}
