package main

import "fmt"

func countDistinctIntegers(nums []int) int {
	theMap := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		theMap[nums[i]] = true
		rn := reverseInt(nums[i])
		nums = append(nums, rn)
		theMap[rn] = true
	}
	return len(theMap)
}

func reverseInt(x int) int {
	answer := 0
	for x > 0 {
		answer = answer*10 + x%10
		x = x / 10
	}
	return answer
}

func main() {
	nums := []int{1, 1, 3, 10, 12, 31}
	fmt.Println(countDistinctIntegers(nums))
}
