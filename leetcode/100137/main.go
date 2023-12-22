package main

import "fmt"

func countSubarrays(nums []int, k int) int64 {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	var answer int64 = 0

	left, right := 0, 0
	maxCount := 0

	for ; right < len(nums); right++ {
		if nums[right] == max {
			maxCount++
		}

		for ; maxCount == k; left++ {
			if nums[left] == max {
				maxCount--
			}
		}
		answer = answer + int64(left)
	}

	return answer
}

func main() {
	nums := []int{1, 3, 3, 2, 1}
	k := 2
	fmt.Println(countSubarrays(nums, k))
}
