package main

import "testing"

func TestCustomWaitGroupBasic(t *testing.T) {
	wg := &CustomWaitGroup{}

	wg.Add(1)
	wg.Done()
	wg.Wait()

	wg.Add(2)
	go func() { wg.Done() }()
	go func() { wg.Done() }()
	wg.Wait()
}

func TestCustomWaitGroupNegativeDelta(t *testing.T) {
	wg := &CustomWaitGroup{}

	defer func() {
		if r := recover(); r == nil {
			t.Error("a negative delta value was expected to cause panic")
		}
	}()

	wg.Add(-1)
}

func TestCustomWaitGroupZeroDelta(t *testing.T) {
	wg := &CustomWaitGroup{}
	wg.Add(0)
}
