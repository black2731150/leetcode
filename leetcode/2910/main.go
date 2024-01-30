package main

import "fmt"

func minGroupsForValidAssignment(nums []int) int {
	theMap := make(map[int]int)
	for _, v := range nums {
		theMap[v]++
	}

	minNum := len(nums)
	for _, v := range theMap {
		if v < minNum {
			minNum = v
		}
	}

	for ; ; minNum-- {
		answer := 0
		for _, v := range theMap {
			if v/minNum < v%minNum {
				answer = 0
				break
			}
			answer += (v + minNum) / (minNum + 1)
		}

		if answer > 0 {
			return answer
		}
	}
}

func main() {
	nums := []int{3, 2, 3, 2, 3}
	fmt.Println(minGroupsForValidAssignment(nums))
}
