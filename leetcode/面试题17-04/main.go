package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	answer := 0
	for i := 0; i < n; i++ {
		answer = answer ^ i
		answer = answer ^ nums[i]
	}
	answer = answer ^ n

	return answer
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}
