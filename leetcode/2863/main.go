package main

import "fmt"

func maxSubarrayLength(nums []int) int {
	answer := 0

	for i := 0; i < len(nums); i++ {
		right := len(nums) - 1
		for i < right {
			if nums[right] < nums[i] {
				answer = max(answer, right-i+1)
				break
			}
			right--
		}
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
	nums := []int{7, 6, 5, 4, 3, 2, 1, 6, 10, 11}
	fmt.Println(maxSubarrayLength(nums))
}
