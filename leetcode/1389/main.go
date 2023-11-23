package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	n := len(index)
	target := make([]int, 0, len(nums))

	for i := 0; i < n; i++ {
		target = append(target[:index[i]+1], target[index[i]:]...)
		target[index[i]] = nums[i]
	}
	return target
}

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	fmt.Println(createTargetArray(nums, index))
}
