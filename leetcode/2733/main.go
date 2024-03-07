package main

import "fmt"

func findNonMinOrMax(nums []int) int {
	min := nums[0]
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	for _, num := range nums {
		if num > min && num < max {
			return num
		}
	}

	return -1
}

func main() {
	nums := []int{3, 2, 1, 5}
	fmt.Println(findNonMinOrMax(nums))
}
