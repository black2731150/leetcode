package main

import "fmt"

func countCornerRectangles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	answer := 0

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			numOfOne := 0
			for c := 0; c < n; c++ {
				if grid[i][c] == 1 && grid[j][c] == 1 {
					numOfOne++
				}
			}
			answer = answer + ((numOfOne-1)*numOfOne)/2
		}
	}
	return answer
}

func main() {
	grid := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	fmt.Println(countCornerRectangles(grid))
}
