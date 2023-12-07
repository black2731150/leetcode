package main

import "fmt"

func sumOfUnique(nums []int) int {
	theMap := make(map[int]int)

	answer := 0
	for _, v := range nums {
		theMap[v]++
	}

	for k, v := range theMap {
		if v == 1 {
			answer += k
		}
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 2}
	fmt.Println(sumOfUnique(nums))
}
