package main

import "fmt"

func setZeroes(matrix [][]int) {
	x := []int{}
	y := []int{}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				x = append(x, i)
				y = append(y, j)
			}
		}
	}

	for i := range x {
		for j := range matrix[x[i]] {
			matrix[x[i]][j] = 0
		}
	}

	for i := range y {
		for j := range matrix {
			matrix[j][y[i]] = 0
		}
	}
}

func main() {
	matirx := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	setZeroes(matirx)
	fmt.Println(matirx)
}
