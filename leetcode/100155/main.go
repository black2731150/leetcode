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

		x := pow(int64(a), int64(b)) % int64(10)
		y := pow(x, int64(c))
		z := y % int64(m)
		if z == int64(target) {
			answer = append(answer, i)
		}
	}

	return answer
}

func pow(x, y int64) int64 {
	var answer int64 = 1
	var i int64 = 0
	for ; i < y; i++ {
		answer = answer * int64(x)
	}
	return answer
}

func main() {
	variables := [][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}
	target := 2
	fmt.Println(getGoodIndices(variables, target))
}
