package main

import "fmt"

func findColumnWidth(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	answer := make([]int, n)
	for i := 0; i < m; i++ {
		for j, v := range grid[i] {
			o := lenght(v)
			if v < 0 {
				o++
			}

			answer[j] = max(answer[j], o)
		}
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lenght(x int) int {
	if x == 0 {
		return 1
	}
	answer := 0
	for x != 0 {
		x = x / 10
		answer++
	}

	return answer
}

func main() {
	grid := [][]int{{1}, {22}, {333}}
	fmt.Println(findColumnWidth(grid))
}
