package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	answer := 0
	end := 0
	maxPosition := 0
	for i := 0; i < n-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			answer++
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
	nums := []int{3, 2, 1}
	fmt.Println(jump(nums))
}
