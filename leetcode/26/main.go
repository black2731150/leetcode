package main

import "fmt"

func removeDuplicates(nums []int) int {
	lenNums := len(nums)
	for i := 0; i+1 < lenNums; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			lenNums--
		}
	}

	return lenNums
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(nums[:n])
}
