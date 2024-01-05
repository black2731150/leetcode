package main

import "fmt"

func canSeePersonsCount(heights []int) []int {

	n := len(heights)
	stack := []int{}
	answer := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		height := heights[i]
		for len(stack) > 0 && stack[len(stack)-1] <= height {
			stack = stack[:len(stack)-1]
			answer[i]++
		}

		if len(stack) > 0 {
			answer[i]++
		}
		stack = append(stack, height)
	}
	return answer
}

func main() {
	heights := []int{10, 6, 8, 5, 11, 9}
	fmt.Println(canSeePersonsCount(heights))
}
