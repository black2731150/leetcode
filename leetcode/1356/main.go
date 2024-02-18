package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		ni := numOfOne(arr[i])
		nj := numOfOne(arr[j])
		if ni == nj {
			return arr[i] < arr[j]
		}

		return ni < nj
	})

	return arr
}

func numOfOne(x int) int {
	answer := 0
	for i := 0; i < 16; i++ {
		if 1&(x>>i) == 1 {
			answer++
		}
	}

	return answer
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortByBits(arr))
}
