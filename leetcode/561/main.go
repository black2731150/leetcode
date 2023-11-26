package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	answer := 0
	n := len(nums)

	for i := 0; i < n; i = i + 2 {
		answer = answer + nums[i]
	}

	return answer
}

func main() {
	nums := []int{1, 2, 2, 4, 5, 6}
	fmt.Println(arrayPairSum(nums))
}
