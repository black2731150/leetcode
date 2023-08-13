package main

import (
	"fmt"
	"sort"
)

// func hIndex(citations []int) int {
// 	lenCitations := len(citations)
// 	answer := 0
// 	for i := lenCitations; i >= 1; i-- {
// 		tmp := 0
// 		for j := range citations {
// 			if citations[j] >= i {
// 				tmp++
// 			}
// 		}
// 		if tmp >= i {
// 			answer = max(answer, i)
// 		}
// 	}
// 	return answer
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

func hIndex(citations []int) int {
	lenCitations := len(citations)
	sort.Ints(citations)
	answer := 0
	for i := lenCitations - 1; i >= 0; i-- {
		if citations[i] > answer {
			answer++
		}
	}

	return answer
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations))
}
