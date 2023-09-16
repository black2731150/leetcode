package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	numsLen := len(nums)
	theMap := make(map[int]int, numsLen)
	for i := range nums {
		if last, find := theMap[nums[i]]; find {
			if i-last <= k {
				return true
			} else {
				theMap[nums[i]] = i
			}
		} else {
			theMap[nums[i]] = i
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
}
