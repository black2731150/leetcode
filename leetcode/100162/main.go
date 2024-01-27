package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	theMap := make(map[int]int)
	max := 0

	for _, v := range nums {
		theMap[v]++
		if x := theMap[v]; x > max {
			max = x
		}
	}

	answer := 0
	for _, v := range theMap {
		if v == max {
			answer += max
		}
	}

	return answer
}

func main() {
	nums := []int{1, 2, 2, 3, 1, 4}
	fmt.Println(maxFrequencyElements(nums))
}
