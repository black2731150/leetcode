package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)

	left := 0
	right := n - 1

	index := -1

	for left <= right {
		if nums[left] == target {
			index = left
			break
		}

		if nums[right] == target {
			index = right
			break
		}

		mid := (left + right) / 2

		if nums[mid] == target {
			index = mid
			break
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if index == -1 {
		return []int{-1, -1}
	}

	leftindex := index
	for i := leftindex; i >= 0; i-- {
		if nums[leftindex] == nums[index] {
			leftindex--
		}
	}

	rightindex := index
	for i := rightindex; i < n; i++ {
		if nums[rightindex] == nums[index] {
			rightindex++
		}
	}

	return []int{leftindex + 1, rightindex - 1}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 5, 6}
	target := 5
	fmt.Println(searchRange(nums, target))
}
