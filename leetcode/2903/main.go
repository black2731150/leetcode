package main

import "fmt"

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	answer := [2]int{-1, -1}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if abs(j-i) >= indexDifference && abs(nums[i]-nums[j]) >= valueDifference {
				answer[0] = i
				answer[1] = j
			}
		}
	}
	return answer[:]
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	nums := []int{5, 1, 4, 1}
	indexDifference := 1
	valueDifference := 2
	fmt.Println(findIndices(nums, indexDifference, valueDifference))
}
