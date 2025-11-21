package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	chanUint8 := make(chan uint8)
	chanFloat64 := make(chan float64)

	go generateUint8(chanUint8, 5)
	go cubing(chanUint8, chanFloat64)

	for v := range chanFloat64 {
		fmt.Println(v)
	}
}

func generateUint8(output chan<- uint8, count int) {
	defer close(output)
	for range count {
		output <- uint8(rand.Intn(256))
	}
}

func cubing(input <-chan uint8, output chan<- float64) {
	defer close(output)
	for v := range input {
		output <- math.Pow(float64(v), 3)
	}
}
