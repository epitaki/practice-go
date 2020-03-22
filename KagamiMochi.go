package main

import (
	"fmt"
	"sort"
)

func kagamiMochi(inputArray []int) int {
	sort.Ints(inputArray)
	j := 0
	for i := 1; i < len(inputArray); i++ {
		if inputArray[j] == inputArray[i] {
			continue
		}
		j++
		inputArray[j] = inputArray[i]
	}
	return j + 1
}

func main() {
	var N int
	var num int
	fmt.Scan(&N)
	inputArray := []int{}
	for i := 0; i < N; i++ {

		fmt.Scan(&num)
		inputArray = append(inputArray, num)
	}
	fmt.Println(kagamiMochi(inputArray))
}
