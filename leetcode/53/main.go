package main

import "fmt"

func maxSubArray(nums []int) int {
	answer := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > answer {
			answer = nums[i]
		}
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(maxSubArray(nums))
}
