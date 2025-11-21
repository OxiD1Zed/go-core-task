package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := CustomWaitGroup{}
	wg.Add(2)

	f := func() {
		for i := range 1000 {
			fmt.Print(i, " ")
		}
		wg.Done()
	}
	go f()
	go f()

	wg.Wait()
}

type CustomWaitGroup struct {
	mu      sync.Mutex
	counter int
	done    chan struct{}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	if delta == 0 {
		return
	}
	if wg.counter+delta < 0 {
		panic("The delta is larger than the counter")
	}

	if wg.counter == 0 {
		wg.done = make(chan struct{})
	}
	wg.counter += delta
	if wg.counter == 0 {
		close(wg.done)
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.done
}
