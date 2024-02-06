package main

import "sort"

func wiggleSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	sort.Ints(nums)

	n := len(nums)

	for i := 1; i < n; i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}

	return
}
