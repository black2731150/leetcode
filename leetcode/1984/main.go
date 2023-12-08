package main

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	sort.Ints(nums)

	answer := 10000000
	for i := 0; i < len(nums); i++ {
		if i+k-1 < len(nums) {
			answer = min(answer, nums[i+k-1]-nums[i])
		}
	}

	return answer
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := []int{9, 4, 1, 7}
	k := 2
	fmt.Println(minimumDifference(nums, k))
}
