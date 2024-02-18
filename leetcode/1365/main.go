package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	cn := make([]int, len(nums))
	copy(cn, nums)
	sort.Ints(cn)
	n := len(nums)

	fmt.Println(cn)
	answer := []int{}
	for i := 0; i < n; i++ {
		v := nums[i]
		left, right := 0, n-1
		for left <= right {
			mid := (left + right) / 2

			if cn[mid] < v {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		answer = append(answer, left)
	}
	return answer
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
