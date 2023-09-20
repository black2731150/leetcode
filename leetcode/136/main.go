package main

import "fmt"

func singleNumber(nums []int) int {
	n := len(nums)
	answer := nums[0]
	for i := 1; i < n; i++ {
		answer = answer ^ nums[i]
	}

	return answer
}

func main() {
	nums := []int{1, 2, 2, 1, 4}
	fmt.Println(singleNumber(nums))
}
