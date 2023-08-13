package main

import "fmt"

func majorityElement(nums []int) int {
	lenNums := len(nums)
	theMap := make(map[int]int, lenNums)
	for i := range nums {
		if theMap[nums[i]]++; theMap[nums[i]] > (lenNums / 2) {
			return nums[i]
		}
	}

	return 0
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
