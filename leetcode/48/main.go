package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-i-1][j]
			matrix[n-i-1][j] = tmp
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {

			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

func main() {
	matirx := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matirx)
	fmt.Println(matirx)
}
