package main

import "fmt"

func shuffle(nums []int, n int) []int {

	answer := make([]int, 2*n)
	ai := 0
	for i := 0; i < n; i++ {
		answer[ai] = nums[i]
		answer[ai+1] = nums[i+n]
		ai += 2
	}
	return answer
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := len(nums) / 2
	fmt.Println(shuffle(nums, n))
}
