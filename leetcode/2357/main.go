package main

import "fmt"

func minimumOperations(nums []int) int {

	answer := 0

	for !isAllZero(nums) {
		minNum := min(nums)
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				continue
			}
			nums[i] = nums[i] - minNum
		}
		answer++
	}
	return answer
}

func isAllZero(nums []int) bool {
	for _, v := range nums {
		if v == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

func min(nums []int) int {
	min := nums[0]
	for i := range nums {
		if nums[i] == 0 {
			continue
		}

		if min == 0 {
			min = nums[i]
		}

		if min != 0 && nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func main() {
	nums := []int{1, 5, 0, 3, 5}
	fmt.Println(minimumOperations(nums))
}
