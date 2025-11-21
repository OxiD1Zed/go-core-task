package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	outInt := make(chan int)
	outInt2 := make(chan int)
	outInt3 := make(chan int)
	mergeInt := make(chan int)

	go generateInt(outInt, 5)
	go generateInt(outInt2, 5)
	go generateInt(outInt3, 5)
	go mergeChannels(mergeInt, outInt, outInt2, outInt3)

	for v := range mergeInt {
		fmt.Println(v)
	}
}

func mergeChannels[T any](output chan<- T, input ...<-chan T) {
	defer close(output)
	wg := sync.WaitGroup{}
	wg.Add(len(input))
	for _, ch := range input {
		go func(ch <-chan T) {
			defer wg.Done()
			for v := range ch {
				output <- v
			}
		}(ch)
	}
	wg.Wait()
}

func generateInt(ch chan<- int, count int) {
	defer close(ch)
	for range count {
		ch <- rand.Int()
	}
}
