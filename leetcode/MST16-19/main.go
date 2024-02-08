package main

import (
	"fmt"
	"sort"
)

func pondSizes(land [][]int) []int {
	m := len(land)
	n := len(land[0])
	answer := []int{}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || y < 0 || x >= m || y >= n || land[x][y] != 0 {
			return 0
		}

		land[x][y] = -1
		return 1 + dfs(x-1, y-1) + dfs(x-1, y) + dfs(x-1, y+1) + dfs(x, y-1) + dfs(x, y+1) + dfs(x+1, y-1) + dfs(x+1, y) + dfs(x+1, y+1)
	}

	for i, v := range land {
		for i2, v2 := range v {
			if v2 == 0 {
				answer = append(answer, dfs(i, i2))
			}
		}
	}

	sort.Ints(answer)

	return answer
}

func main() {
	land := [][]int{
		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}

	fmt.Println(pondSizes(land))
}
