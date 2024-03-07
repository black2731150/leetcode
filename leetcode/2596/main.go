package main

import "fmt"

func checkValidGrid(grid [][]int) bool {
	n := len(grid)

	last := -1

	p := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, -2}, {1, 2}, {-1, 2}, {-1, -2}}
	var dfs func(x, y int, target int)
	dfs = func(x, y int, target int) {
		find := false
		nextx, nexty := 0, 0
		for _, pp := range p {
			nx := x + pp[0]
			ny := y + pp[1]

			if check(nx, ny, n) && grid[nx][ny] == target {
				find = true
				nextx = nx
				nexty = ny
				break
			}
		}

		if find {
			dfs(nextx, nexty, target+1)
		} else {
			last = target - 1
		}
	}

	if grid[0][0] != 0 {
		return false
	} else {
		dfs(0, 0, 1)
	}

	return last == n*n-1
}

func check(x, y int, n int) bool {
	if x < 0 || y < 0 || x >= n || y >= n {
		return false
	}
	return true
}

func main() {
	grid := [][]int{
		{0, 11, 16, 5, 20},
		{17, 4, 19, 10, 15},
		{12, 1, 8, 21, 6},
		{3, 18, 23, 14, 9},
		{24, 13, 2, 7, 22},
	}
	fmt.Println(checkValidGrid(grid))
}
