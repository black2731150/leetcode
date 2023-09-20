package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)

	left := 0
	right := n - 1

	answer := -1

	for left <= right {
		if nums[left] == target {
			return left
		}

		if nums[right] == target {
			return right
		}

		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[0] < nums[mid] {
			if nums[0] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return answer
}

func main() {
	nums := []int{8, 1, 2, 3, 4, 5, 6, 7}
	target := 6
	fmt.Println(search(nums, target))
}
