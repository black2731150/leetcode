package main

import "fmt"

func missingNumber(nums []int) int {
	hasNum := make(map[int]bool, len(nums)+1)
	// for i := 0; i < len(nums)+1; i++ {
	// 	hasNum[i] = false
	// }

	for i := 0; i < len(nums); i++ {
		hasNum[nums[i]] = true
	}

	for i := 0; i < len(nums)+1; i++ {
		if !hasNum[i] {
			return i
		}
	}

	return 0
}

func main() {
	nums := []int{0, 1, 3}
	fmt.Println(missingNumber(nums))
}
