package main

import "fmt"

func countElements(nums []int) int {
	answer := 0

	max := nums[0]
	min := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	for _, v := range nums {
		if v > min && v < max {
			answer++
		}
	}

	return answer
}

func main() {
	nums := []int{-3, -3 - 3, 90}
	fmt.Println(countElements(nums))
}
