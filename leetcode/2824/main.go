package main

import "fmt"

func countPairs(nums []int, target int) int {
	n := len(nums)
	answer := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] < target {
				answer++
			}
		}
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 3
	fmt.Println(countPairs(nums, target))
}
