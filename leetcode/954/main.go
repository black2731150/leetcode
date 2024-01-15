package main

import (
	"fmt"
	"sort"
)

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)

	zeroNum := 0
	theMap := make(map[int]int)
	for _, v := range arr {
		if v == 0 {
			zeroNum++
		}
		theMap[v]++
	}

	answer := len(arr)
	for _, v := range arr {
		if theMap[v] > 0 && theMap[2*v] > 0 {
			theMap[2*v]--
			theMap[v]--
			answer -= 2
		}
	}

	return answer == 0 && zeroNum%2 == 0
}

func main() {
	arr := []int{2, 4, 0, 0, 8, 1}
	fmt.Println(canReorderDoubled(arr))
}
