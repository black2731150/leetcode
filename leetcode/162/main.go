package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	n := len(nums)

	left := 0
	right := n - 1

	bian := func(x int) int {
		if x == -1 || x == n {
			return math.MinInt
		}
		return nums[x]
	}

	for {
		mid := (left + right) / 2

		if bian(mid) > bian(mid-1) && bian(mid) > bian(mid+1) {
			return mid
		}

		if bian(mid) >= bian(mid+1) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(findPeakElement(nums))
}
