package main

import (
	"fmt"
	"strconv"
)

func lastVisitedIntegers(words []string) []int {
	stack := []int{}
	answer := []int{}
	last := 0
	for _, word := range words {
		if word == "prev" {
			if last >= 0 && len(stack) > 0 {
				answer = append(answer, stack[last])
			} else {
				answer = append(answer, -1)
			}
			last--
		} else {
			num, _ := strconv.Atoi(word)
			stack = append(stack, num)
			last = len(stack) - 1
		}
	}

	return answer
}

func main() {
	words := []string{"1", "2", "prev", "prev", "prev"}
	fmt.Println(lastVisitedIntegers(words))
}
