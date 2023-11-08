package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func main() {
	nums := []int{1, 4, 6, 7, 8}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
