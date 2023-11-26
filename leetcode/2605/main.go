package main

import (
	"fmt"
	"sort"
)

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				return nums1[i]
			}
		}
	}

	if nums1[0] < nums2[0] {
		return nums1[0]*10 + nums2[0]
	} else {
		return nums2[0]*10 + nums1[0]
	}
}

func main() {
	nums1 := []int{1, 2, 4}
	nums2 := []int{4, 5, 7}
	fmt.Println(minNumber(nums1, nums2))
}
