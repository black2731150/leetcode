package main

import "fmt"

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition(nums, low, high)
		quickSort(nums, low, pivotIndex-1)
		quickSort(nums, pivotIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	i, j := low, high

	for i < j {
		for i < j && nums[i] < pivot {
			i++
		}

		for j > i && nums[j] > pivot {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return j
}

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortArray(nums))
}
