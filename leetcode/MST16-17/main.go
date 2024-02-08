package main

import "fmt"

func maxSubArray(nums []int) int {
	answer := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}

		answer = max(answer, nums[i])
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
