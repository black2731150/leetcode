package main

import "fmt"

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m := len(mat1)
	k := len(mat2)
	n := len(mat2[0])

	answer := make([][]int, m)
	for i := range answer {
		answer[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for z := 0; z < k; z++ {
				answer[i][j] += mat1[i][z] * mat2[z][j]
			}
		}
	}

	return answer
}

func main() {
	mat1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	mat2 := [][]int{{2, 3, 4}, {5, 6, 7}}
	fmt.Println(multiply(mat1, mat2))
}
