package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
	answer := 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[i])-3; j++ {
			if help(grid, i, j) {
				answer++
			}
		}
	}
	return answer
}

func help(gird [][]int, i, j int) bool {

	theMap := make(map[int]struct{})
	for i1 := i; i1 < i+3; i1++ {
		for j1 := j; j1 < j+3; j1++ {
			if gird[i1][j1] < 1 || gird[i1][j1] >= 10 {
				return false
			}
			theMap[gird[i1][j1]] = struct{}{}
		}
	}

	if len(theMap) != 9 {
		return false
	}

	a1, a2, a3 := gird[i+0][j+0], gird[i+0][j+1], gird[i+0][j+2]
	b1, b2, b3 := gird[i+1][j+0], gird[i+1][j+1], gird[i+1][j+2]
	c1, c2, c3 := gird[i+2][j+0], gird[i+2][j+1], gird[i+2][j+2]

	x := a1 + a2 + a3

	if b1+b2+b3 != x || c1+c2+c3 != x || a1+b1+c1 != x || a2+b2+c2 != x || a3+b3+c3 != x || a1+b2+c3 != x || a3+b2+c1 != x {
		return false
	}

	return true
}

func main() {
	gird := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
	fmt.Println(numMagicSquaresInside(gird))
}
