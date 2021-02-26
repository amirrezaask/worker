package worker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFlow(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	w := NewSimpleWorker().
		WithQueue(NewChannelQueue(10)).
		WithNGouroutines(5)
	w.Start()
	w.RunJob(NewSimpleJob("sample job", func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("salam man ra beshenavid")
		wg.Done()
		return nil
	}))
	wg.Wait()
}
