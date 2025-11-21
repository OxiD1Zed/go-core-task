package main

import (
	"slices"
	"testing"
)

func TestIntersectionSlices(t *testing.T) {
	a := []int{65, 3, 58, 678, 64, 64}
	b := []int{64, 64, 2, 3, 43}
	received_bool, received := intersectionSlices(a, b)

	expected := []int{64, 64, 3}
	expected_bool := true

	if slices.Compare(expected, received) != 0 && expected_bool == received_bool {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}
