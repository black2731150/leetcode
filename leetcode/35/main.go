package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := len(nums)

	left := 0
	right := n - 1

	answer := n

	for left <= right {

		if nums[left] == target {
			return left
		}

		if nums[right] == target {
			return right
		}

		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
			answer = mid
		} else {
			left = mid + 1
		}
	}

	return answer
}

func main() {
	nums := []int{1, 3, 4, 5}
	fmt.Println(searchInsert(nums, 2))
}
