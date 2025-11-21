package main

import (
	"maps"
	"testing"
)

func TestAdd(t *testing.T) {
	m := StringIntMap{}
	m.Add("s", 1)
	received := len(m)

	expected := 1

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestRemove(t *testing.T) {
	m := StringIntMap{}
	m.Add("s", 1)
	startLen := len(m)
	m.Remove("s")
	endLen := len(m)
	received := startLen > endLen

	expected := true

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestCopy(t *testing.T) {
	m := StringIntMap{}
	m.Add("s", 1)
	copyM := m.Copy()
	received := maps.Equal(m, copyM)

	expected := true

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestExists(t *testing.T) {
	m := StringIntMap{}
	m.Add("s", 1)
	received := m.Exists("s")

	expected := true

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}

func TestGet(t *testing.T) {
	m := StringIntMap{}
	_, received := m.Get("m")

	expected := false

	if expected != received {
		t.Errorf("expected: %v, received: %v", expected, received)
	}
}
