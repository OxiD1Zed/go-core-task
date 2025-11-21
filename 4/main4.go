package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(differenceSlices(slice1, slice2))
}

func differenceSlices(sl1, sl2 []string) []string {
	elemenetsMap := map[string]bool{}

	for _, value := range sl2 {
		elemenetsMap[value] = true
	}

	res := []string{}

	for _, value := range sl1 {
		if !elemenetsMap[value] {
			res = append(res, value)
		}
	}

	return res
}
