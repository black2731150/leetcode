package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	if len(nums)%3 != 0 {
		return [][]int{}
	}

	answer := [][]int{}

	for i := 0; i < len(nums); i = i + 3 {
		if abs(nums[i]-nums[i+1]) <= k && abs(nums[i]-nums[i+2]) <= k && abs(nums[i+1]-nums[i+2]) <= k {
			answer = append(answer, []int{nums[i], nums[i+1], nums[i+2]})
		} else {
			return [][]int{}
		}
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 2
	fmt.Println(divideArray(nums, k))
}
