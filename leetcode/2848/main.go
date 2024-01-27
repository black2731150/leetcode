package main

import "fmt"

func numberOfPoints(nums [][]int) int {
	gets := make([]bool, 101)
	for _, v := range nums {
		start := v[0]
		end := v[1]
		for i := start; i <= end; i++ {
			gets[i] = true
		}
	}

	answer := 0
	for _, v := range gets {
		if v {
			answer++
		}
	}
	return answer
}

// func numberOfPoints(nums [][]int) int {
// 	theMap := make(map[int]struct{})
// 	for _, v := range nums {
// 		start := v[0]
// 		end := v[1]
// 		for i := start; i <= end; i++ {
// 			theMap[i] = struct{}{}
// 		}
// 	}
// 	return len(theMap)
// }

func main() {
	nums := [][]int{{3, 6}, {1, 5}, {4, 7}}
	fmt.Println(numberOfPoints(nums))
}
