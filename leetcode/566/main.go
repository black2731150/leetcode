package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	// m := len(mat)
	n := len(mat[0])

	answer := make([][]int, r)
	for x := range answer {
		answer[x] = make([]int, c)
	}

	i, j := 0, 0
	for _, row := range answer {
		for y := range row {
			row[y] = mat[i][j]
			j++
			if j == n {
				j = 0
				i++
			}
		}
	}
	return answer
}

func main() {
	mat := [][]int{{1, 2}, {3, 4}}
	r := 1
	c := 4
	fmt.Println(matrixReshape(mat, r, c))
}
