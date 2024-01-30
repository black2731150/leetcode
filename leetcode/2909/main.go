package main

import (
	"fmt"
	"math"
)

func minimumSum(nums []int) int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)

	left[0] = nums[0]
	for i := 1; i < n; i++ {
		left[i] = min(left[i-1], nums[i])
	}

	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = min(right[i+1], nums[i])
	}

	answer := math.MaxInt
	for j := 1; j < n-1; j++ {
		if nums[j] > left[j] && nums[j] > right[j] {
			if x := left[j] + right[j] + nums[j]; x < answer {
				answer = x
			}
		}
	}

	if answer == math.MaxInt {
		return -1
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
	nums := []int{6, 5, 4, 3, 4, 5}
	fmt.Println(minimumSum(nums))
}
