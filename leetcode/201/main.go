package main

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	answer := left

	for i := left; i <= right; i++ {
		answer = answer & i
		if answer == 0 {
			return 0
		}
	}

	return answer
}

func main() {
	left := 4
	right := 7

	fmt.Println(rangeBitwiseAnd(left, right))
}
