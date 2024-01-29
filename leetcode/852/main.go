package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)
	for left <= right {
		mid := (left + right) / 2
		if mid == 0 && mid == len(arr)-1 {
			return mid
		}

		if arr[mid]-arr[mid-1] > 0 && arr[mid+1]-arr[mid] > 0 {
			left = mid
		}

		if arr[mid]-arr[mid-1] < 0 && arr[mid+1]-arr[mid] < 0 {
			right = mid
		}

		if arr[mid-1] < arr[mid] && arr[mid+1] < arr[mid] {
			return mid
		}
	}

	return 0
}

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}
