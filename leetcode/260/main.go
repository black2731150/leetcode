package main

import "fmt"

func singleNumber(nums []int) []int {
	n := len(nums)

	theMap := make(map[int]bool, n)

	answer := []int{}

	for i := range nums {
		theMap[nums[i]] = !theMap[nums[i]]
	}

	for i := range theMap {
		if theMap[i] == true {
			answer = append(answer, i)
		}
	}

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(singleNumber(nums))
}
