package main

import "fmt"

func firstMissingPositive(nums []int) int {
	numMap := make(map[int]struct{})

	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	for i := 1; ; i++ {
		if _, find := numMap[i]; !find {
			return i
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(firstMissingPositive(nums))
}
