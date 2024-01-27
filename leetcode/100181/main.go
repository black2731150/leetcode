package main

import (
	"fmt"
	"sort"
)

func minimumCost(nums []int) int {
	first := nums[0]

	nums = nums[1:]
	sort.Ints(nums)

	return first + nums[0] + nums[1]
}

func main() {
	nums := []int{10, 2, 4, 8}
	fmt.Println(minimumCost(nums))
}
