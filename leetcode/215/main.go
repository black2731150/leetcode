package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	qSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func qSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		qSort(arr, low, p-1)
		qSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	nums := []int{1, 4, 6, 7, 8}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
