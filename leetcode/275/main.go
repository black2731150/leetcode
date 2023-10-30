package main

import (
	"fmt"
)

func hIndex(citations []int) int {
	n := len(citations)

	left := 0
	right := n - 1

	answer := 0
	for left <= right {
		mid := (left + right) / 2
		if citations[mid] >= n-mid {
			answer = answer + (right - mid + 1)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return answer
}

// func hIndex(citations []int) int {
// 	n := len(citations)
// 	return n - sort.Search(n, func(x int) bool { return citations[x] >= n-x })
// }

func main() {
	citations := []int{1}
	fmt.Println(hIndex(citations))
}
