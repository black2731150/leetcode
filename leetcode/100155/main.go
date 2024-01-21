package main

import (
	"fmt"
)

func getGoodIndices(variables [][]int, target int) []int {
	answer := []int{}

	for i, nums := range variables {
		a := nums[0]
		b := nums[1]
		c := nums[2]
		m := nums[3]

		x := powMod(int64(a), int64(b), 10)
		y := powMod(x, int64(c), int64(m))
		if y == int64(target) {
			answer = append(answer, i)
		}
	}

	return answer
}

func powMod(x, y, mod int64) int64 {
	var answer int64 = 1
	for i := int64(0); i < y; i++ {
		answer = (answer * x) % mod
	}
	return answer
}

func main() {
	variables := [][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}
	target := 2
	fmt.Println(getGoodIndices(variables, target))
}
