package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	lenNums := len(nums)
	k = k % lenNums
	first := nums[lenNums-k:]
	last := nums[:lenNums-k]
	first = append(first, last...)
	copy(nums, first)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
