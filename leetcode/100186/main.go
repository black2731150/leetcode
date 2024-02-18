package main

import "fmt"

func countMatchingSubarrays(nums []int, pattern []int) int {
	m := len(pattern)
	n := len(nums)

	answer := 0
	for i := 0; i < n-m; i++ {
		ok := true
		for k := 0; k < m; k++ {
			switch pattern[k] {
			case 1:
				if nums[i+k+1] > nums[i+k] {
					continue
				} else {
					ok = false
				}

			case 0:
				if nums[i+k+1] == nums[i+k] {
					continue
				} else {
					ok = false
				}

			case -1:
				if nums[i+k+1] < nums[i+k] {
					continue
				} else {
					ok = false
				}
			}
		}

		if ok {
			answer++
		}
	}

	return answer
}

func main() {
	nums := []int{3, 5, 6, 7, 2}
	pattern := []int{1, 1, 1, -1, 1}
	fmt.Println(countMatchingSubarrays(nums, pattern))
}
