package main

import "fmt"

type Location struct {
	x int
	y int
}

func spiralOrder(matrix [][]int) []int {

	Fangxiang := []Location{
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: 0},
	}
	findex := 0

	m := len(matrix)
	n := len(matrix[0])
	answer := make([]int, 0)

	haveMap := make([][]bool, m)
	for i := range haveMap {
		haveMap[i] = make([]bool, n)
	}

	Location := Fangxiang[findex]
	end := 0
	for x, y := 0, 0; ; {
		if (x >= 0 && x <= m-1) && (y >= 0 && y <= n-1) && !haveMap[x][y] {
			answer = append(answer, matrix[x][y])
			haveMap[x][y] = true
			end = 0
		}

		if (x+Location.x >= 0 && x+Location.x <= m-1) && (y+Location.y >= 0 && y+Location.y <= n-1) && !haveMap[x+Location.x][y+Location.y] {
			x = x + Location.x
			y = y + Location.y
		} else {
			findex = findex + 1
			Location = Fangxiang[findex%4]
			end++
			if end == 5 {
				break
			}
		}
	}

	return answer
}

func main() {
	matirx := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(matirx))
}
