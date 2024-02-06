package main

import (
	"fmt"
)

func divingBoard(shorter int, longer int, k int) []int {
	answer := make([]int, k+1)
	if k == 0 {
		return []int{}
	}

	if shorter == longer {
		return []int{k * longer}
	}
	for i := 0; i <= k; i++ {
		answer[i] = i*longer + shorter*(k-i)
	}
	// sort.Ints(answer)
	return answer
}

func main() {
	shorter := 1
	longer := 2
	k := 3
	fmt.Println(divingBoard(shorter, longer, k))
}
