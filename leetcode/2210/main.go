package main

import "fmt"

// func countHillValley(nums []int) int {
// 	answer := 0
// 	nums := []int{nums[0]}

// 	for i := 1; i < len(nums); i++ {
// 		if nums[len(nums)-1] == nums[i] {
// 			continue
// 		} else {
// 			nums = append(nums, nums[i])
// 		}
// 	}

// 	for i := 1; i < len(nums)-1; {
// 		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
// 			answer++
// 			i++
// 			continue
// 		}

// 		if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
// 			answer++
// 			i++
// 			continue
// 		}
// 		i++
// 	}
// 	return answer
// }

func countHillValley(nums []int) int {
	answer := 0

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			answer++
			i++
			continue
		}

		if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			answer++
			i++
			continue
		}

		for nums[i] == nums[i+1] {
			i++
		}
	}
}

func main() {
	nums := []int{2, 4, 1, 1, 6, 5}
	fmt.Println(countHillValley(nums))
}
