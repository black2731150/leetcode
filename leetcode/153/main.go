package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums)

	left := 0
	right := n - 1

	min := 5001

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] < min {
			min = nums[mid]
		}

		if nums[0] <= nums[mid] {
			left = mid + 1
			if nums[0] < min {
				min = nums[0]
			}
		} else {
			right = mid - 1
			if nums[mid] < min {
				min = nums[mid]
			}
		}
	}
	return min
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
