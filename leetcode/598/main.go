package main

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	minA, minB := 500000, 50000
	for _, v := range ops {
		minA = min(minA, v[0])
		minB = min(minB, v[1])
	}

	return minA * minB
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	m := 3
	n := 2
	ops := [][]int{{1, 1}}
	fmt.Println(maxCount(m, n, ops))
}
