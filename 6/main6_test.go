package main

import (
	"slices"
	"testing"
)

func TestGenerateInt(t *testing.T) {
	ch := make(chan int)
	go generateInt(ch, 10)
	sl1 := make([]int, 5)
	sl2 := make([]int, 5)
	for i := range 5 {
		sl1[i] = <-ch
		sl2[i] = <-ch
	}
	received := slices.Compare(sl1, sl2) != 0

	expected := true

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}
