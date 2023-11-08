package main

import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)

	answer := 0

	min := nums[0]
	max := nums[0]
	for i := 0; i < n; i++ {
		if nums[i]-min <= k {
			max = nums[i]
		} else {
			answer++
			min = nums[i]
			max = nums[i]
		}
	}
	max = max + 1
	return answer + 1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 2
	fmt.Println(partitionArray(nums, k))
}
