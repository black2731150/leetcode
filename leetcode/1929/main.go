package main

import "fmt"

func getConcatenation(nums []int) []int {
	answer := make([]int, 2*len(nums))
	for i, v := range nums {
		answer[i] = v
		answer[i+len(nums)] = v
	}
	return answer
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(getConcatenation(nums))
}
