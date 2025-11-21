package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sl := originalSlice()
	fmt.Println("originSlice: ", sl)
	slEX := sliceExample(sl)
	fmt.Println("sliceExample: ", slEX)
	addEl := addElements(slEX, 151)
	fmt.Println("addElement: ", addEl)
	copySl := copySlice(addEl)
	fmt.Println("copySlice: ", copySl)
	copySl[1] = 1222
	fmt.Println("change the element with index 1 to 122")
	fmt.Println("the original segment before copying: ", addEl, "copied segment: ", copySl)
	remSl := removeElement(copySl, 0)
	fmt.Println("removeElement: ", remSl)
}

func originalSlice() []int {
	res := make([]int, 10)
	for i := range res {
		res[i] = rand.Int()
	}

	return res
}

func sliceExample(sl []int) []int {
	res := []int{}
	index := 0
	for _, v := range sl {
		if v%2 == 0 {
			res = append(res, v)
			index++
		}
	}

	return res
}

func addElements(sl []int, value int) []int {
	return append(sl, value)
}

func copySlice(sl []int) []int {
	return append([]int{}, sl...)
}

func removeElement(sl []int, index int) []int {
	return append(sl[:index], sl[index+1:]...)
}
