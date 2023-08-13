package main

import "fmt"

func removeDuplicates(nums []int) int {
	lenNums := len(nums)
	for i := 0; i+2 < lenNums; i++ {
		if nums[i] == nums[i+1] && nums[i+1] == nums[i+2] {
			nums = append(nums[:i+1], nums[i+2:]...)
			lenNums--
			i--
		}
	}

	return lenNums
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	n := removeDuplicates(nums)
	fmt.Println(nums[:n])
}
