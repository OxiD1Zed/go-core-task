package main

import (
	"slices"
	"testing"
)

func TestGenerateUint8(t *testing.T) {
	chUInt8 := make(chan uint8)
	sl1 := make([]uint8, 5)
	sl2 := make([]uint8, 5)
	go generateUint8(chUInt8, 10)
	for i := range 5 {
		sl1[i] = <-chUInt8
		sl2[i] = <-chUInt8
	}
	received := slices.Compare(sl1, sl2) != 0

	expected := true

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestCubing(t *testing.T) {
	ch := make(chan uint8)
	resCh := make(chan float64)
	go func() {
		defer close(ch)
		for i := range 3 {
			ch <- uint8(i)
		}
	}()
	go cubing(ch, resCh)
	received := []float64{}
	for v := range resCh {
		received = append(received, v)
	}

	expected := []float64{0, 1, 8}

	if slices.Compare(expected, received) != 0 {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}
