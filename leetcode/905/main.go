package main

import "fmt"

// 使用了额外空间
// func sortArrayByParity(nums []int) []int {
// 	n := len(nums)
// 	answer := make([]int, len(nums))
// 	left, right := 0, n-1

// 	for _, v := range nums {
// 		if v%2 == 0 {
// 			answer[left] = v
// 			left++
// 		} else {
// 			answer[right] = v
// 			right--
// 		}
// 	}
// 	return answer
// }

// 不使用额外空间
func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		for nums[left]%2 == 0 && left < right {
			left++
		}

		for nums[right]%2 != 0 && left < right {
			right--
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	return nums
}

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}
