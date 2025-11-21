package main

import "testing"

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	resCh := make(chan int)
	go func() {
		for i := range 5 {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := range 5 {
			ch2 <- i + 5
		}
		close(ch2)
	}()
	go mergeChannels(resCh, ch1, ch2)
	resSl := []int{}
	for v := range resCh {
		resSl = append(resSl, v)
	}
	received := len(resSl)

	expected := 10

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}
