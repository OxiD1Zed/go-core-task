package main

import (
	"slices"
	"testing"
)

func TestDifferenceSlices(t *testing.T) {
	sl1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	sl2 := []string{"banana", "date", "fig"}
	received := differenceSlices(sl1, sl2)

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	if slices.Compare(expected, received) != 0 {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}
