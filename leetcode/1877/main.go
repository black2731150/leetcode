package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	answer := nums[len(nums)-1] + nums[0]
	for i := 0; i < n/2; i = i + 1 {
		answer = max(answer, nums[n-i-1]+nums[i])
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(minPairSum(nums))
}
