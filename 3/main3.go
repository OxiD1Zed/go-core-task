package main

import (
	"fmt"
	"maps"
)

func main() {
	m := StringIntMap{}

	m.Add("id", 1)
	fmt.Println("add: ", m)
	m.Remove("id")
	fmt.Println("remove: ", m)
	copy := StringIntMap(m.Copy())
	fmt.Println("copy: ", copy)
	copy.Add("copy", 123)
	fmt.Println("original: ", m, " copy: ", copy)
	fmt.Println("exists key:copy in origin: ", m.Exists("copy"), "in copy: ", copy.Exists("copy"))
	value, ok := copy.Get("copy")
	fmt.Println("get: ", value, ok)
}

type StringIntMap map[string]int

func (m StringIntMap) Add(key string, value int) {
	m[key] = value
}

func (m StringIntMap) Remove(key string) {
	delete(m, key)
}

func (m StringIntMap) Copy() map[string]int {
	res := make(map[string]int, len(m))
	maps.Copy(res, m)

	return res
}

func (m StringIntMap) Exists(key string) bool {
	_, ok := m[key]
	return ok
}

func (m StringIntMap) Get(key string) (int, bool) {
	value, ok := m[key]
	return value, ok
}
