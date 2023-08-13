package main

import (
	"fmt"
	"sort"
)

func minSubArrayLen(target int, nums []int) int {
	numsLen := len(nums)
	sumOfIndex := make([]int, numsLen+1)
	sumOfIndex[0] = 0
	for i := 1; i <= numsLen; i++ {
		sumOfIndex[i] = nums[i-1] + sumOfIndex[i-1]
	}

	var copyNums []int = make([]int, numsLen)
	copy(copyNums, nums)
	sort.Ints(copyNums)
	sumOfMax := make([]int, numsLen+1)
	J := 1
	for i := numsLen - 1; i >= 0; i-- {
		sumOfMax[J] = copyNums[i] + sumOfMax[J-1]
		J++
	}

	for i := 1; i <= numsLen; i++ {
		if sumOfMax[i] < target {
			continue
		}
		start := 0
		end := start + i
		for end <= numsLen {
			if sumOfIndex[end]-sumOfIndex[start] >= target {
				return i
			} else {
				start++
				end = start + i
			}
		}
	}

	return 0
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 15
	fmt.Println(minSubArrayLen(target, nums))
}
