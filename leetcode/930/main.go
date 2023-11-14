package main

import (
	"fmt"
)

func numSubarraysWithSum(nums []int, goal int) int {
	sumCount := make(map[int]int)
	sumCount[0] = 1 // 初始化为0的前缀和出现1次
	sum := 0
	answer := 0

	for _, num := range nums {
		sum += num
		answer += sumCount[sum-goal] // 查找满足条件的前缀和出现次数
		sumCount[sum]++              // 更新当前前缀和出现次数
	}

	return answer
}

func main() {
	nums := []int{1, 0, 1, 0, 1, 0}
	goal := 2
	fmt.Println(numSubarraysWithSum(nums, goal))
}

// func numSubarraysWithSum(nums []int, goal int) int {
// 	n := len(nums)

// 	sums := make([]int, n)

// 	answer := 0
// 	for i := range nums {
// 		sums = append(sums, sums[i-1]+nums[i])
// 	}

// 	return answer
// }
