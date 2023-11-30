package main

import "fmt"

func expectNumber(scores []int) int {
	theMap := make(map[int]int)

	for _, v := range scores {
		theMap[v]++
	}

	answer := 0
	// 对于每个不同的能力值，期望的匹配次数是 1
	for range theMap {
		answer++
	}

	return answer
}

func main() {
	score := []int{1, 2, 3, 4}
	fmt.Println(expectNumber(score))
}
