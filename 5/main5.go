package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64, 64}
	b := []int{64, 64, 2, 3, 43}

	fmt.Println(intersectionSlices(a, b))
}

func intersectionSlices(sl1, sl2 []int) (bool, []int) {
	elementsMap := map[int]int{}

	for _, value := range sl1 {
		elementsMap[value]++
	}

	res := []int{}

	for _, value := range sl2 {
		if count, ok := elementsMap[value]; ok && count != 0 {
			elementsMap[value]--
			res = append(res, value)
		}
	}

	return len(res) > 0, res
}
