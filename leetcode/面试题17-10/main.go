package main

import "fmt"

func majorityElement(nums []int) int {
	theMap := make(map[int]int)
	answer := -1
	for _, v := range nums {
		theMap[v]++
	}

	for i, v := range theMap {
		if v > len(nums)/2 {
			return i
		}
	}
	return answer
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
