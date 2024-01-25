package main

import "fmt"

func buildArray(nums []int) []int {
	answer := make([]int, len(nums))
	for i := range nums {
		answer[i] = nums[nums[i]]
	}
	return answer
}

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(nums))
}
