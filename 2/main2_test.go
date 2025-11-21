package main

import (
	"slices"
	"testing"
)

func TestOriginalSlice(t *testing.T) {
	sl1 := originalSlice()
	sl2 := originalSlice()
	received := slices.Compare(sl1, sl2) == 0

	expected := false

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestSliceExample(t *testing.T) {
	sl := []int{10, 15, 29, 20, 4}
	received := sliceExample(sl)

	expected := []int{10, 20, 4}

	if slices.Compare(expected, received) != 0 {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestAddElements(t *testing.T) {
	sl := []int{1, 2, 3}
	received := addElements(sl, 11)

	expected := []int{1, 2, 3, 11}

	if slices.Compare(expected, received) != 0 {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestCopySlice(t *testing.T) {
	sl := []int{1, 2, 3}
	received := copySlice(sl)

	expected := []int{1, 2, 3}

	if slices.Compare(expected, received) != 0 {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestDependenciesCopySlice(t *testing.T) {
	sl := []int{1, 2, 3}
	copySl := copySlice(sl)
	copySl[1] = 4
	received := slices.Compare(sl, copySl) == 0

	expected := false

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestRemoveElement(t *testing.T) {
	sl := []int{1, 2, 3}
	received := removeElement(sl, 1)

	expected := []int{1, 3}

	if slices.Compare(expected, received) != 0 {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}
