package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	n := len(nums)
	min := math.MaxInt32
	minNum := 0
	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1

		for left < right {

			if x := abs(target - (nums[i] + nums[left] + nums[right])); x < min {
				min = x
				minNum = (nums[i] + nums[left] + nums[right])
			}

			if nums[i]+nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}

		}
	}

	return minNum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{-1, 2, 1, 4}
	fmt.Println(threeSumClosest(nums, 1))
}
