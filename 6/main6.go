package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)

	go generateInt(ch, 5)

	for v := range ch {
		fmt.Println(v)
	}
}

func generateInt(ch chan<- int, count int) {
	defer close(ch)
	for range count {
		ch <- rand.Int()
	}
}
