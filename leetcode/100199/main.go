package main

import "fmt"

func canSortArray(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				if numOfOne(nums[i]) != numOfOne(nums[j]) {
					return false
				}
			}
		}
	}

	return true
}

func numOfOne(x int) int {
	answer := 0
	for i := 0; i <= 8; i++ {
		if (x>>i)&1 == 1 {
			answer++
		}
	}
	return answer
}

func main() {
	nums := []int{1, 21, 4, 2}
	fmt.Println(canSortArray(nums))
}
