package main

import "fmt"

func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	answers := make([][]int, m)
	for i := range answers {
		answers[i] = make([]int, n)
	}

	answers[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 || i == 0 {
				if j == 0 && i > 0 {
					answers[i][j] = answers[i-1][j] + grid[i][j]
				}

				if i == 0 && j > 0 {
					answers[i][j] = answers[i][j-1] + grid[i][j]
				}
			} else {
				answers[i][j] = min(answers[i-1][j], answers[i][j-1]) + grid[i][j]
			}
		}
	}
	return answers[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
