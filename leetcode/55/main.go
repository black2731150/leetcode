package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	convert := 0

	for i := 0; i <= convert; i++ {
		convert = max(convert, i+nums[i])
		if convert >= n-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
