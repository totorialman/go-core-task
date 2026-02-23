package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type WaitGroup struct {
	counter int64
	sema    chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sema: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(count int) {
	val := atomic.AddInt64(&wg.counter, int64(count))

	if val < 0 {
		panic("WG counter < 0")
	}

	if val == 0 {
		close(wg.sema)
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	if atomic.LoadInt64(&wg.counter) == 0 {
		return
	}

	<-wg.sema
}

func main() {
	wg := NewWaitGroup()
	nWorker := 50

	wg.Add(nWorker)
	for i := 0; i < nWorker; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("worker %d start\n", i)
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
			fmt.Printf("worker %d finish\n", i)
		}()
	}

	fmt.Println("waiting worker")
	wg.Wait()
	fmt.Println("all worker finish")

}
