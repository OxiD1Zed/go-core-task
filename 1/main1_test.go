package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestTypeOfInt8(t *testing.T) {
	var i int8 = 22
	received := TypeOf(i).Kind()

	expected := reflect.Int8

	if expected != received {
		t.Errorf("expected: %s, received: %s", expected, received)
	}
}

func TestContainToString(t *testing.T) {
	a, b := 2, "abc"
	received := ContainToString(a, b)

	expected := "2abc"

	if expected != received {
		t.Errorf("expected: %s, received: %s", expected, received)
	}
}

func TestToSliceRune(t *testing.T) {
	str := "cdc"
	received := ToSliceRune(str)

	expected := []rune{52, 50, 32, 52, 50, 32, 52, 50, 32, 51, 46, 49, 52, 71, 111, 108, 97, 110, 103, 116, 114, 117, 101, 32, 40, 49, 43, 50, 105, 41}

	if slices.Compare(expected, received) == 0 {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestSha256EndcodingWithSalt(t *testing.T) {
	str := "aaa"
	received := Sha256EndcodingWithSalt([]rune(str), []rune("go-2024"))

	expected := [32]byte{231, 22, 119, 18, 224, 20, 53, 89, 111, 79, 24, 43, 106, 83, 116, 1, 129, 213, 167, 115, 222, 65, 69, 253, 81, 188, 100, 225, 164, 224, 177, 247}

	if expected != received {
		t.Errorf("expected: %s, received: %s", expected, received)
	}
}
