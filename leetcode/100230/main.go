package main

import "fmt"

func modifiedMatrix(matrix [][]int) [][]int {
	lieMax := make([]int, len(matrix[0]))

	for _, v := range matrix {
		for j, v2 := range v {
			lieMax[j] = max(lieMax[j], v2)
		}
	}

	answer := make([][]int, len(matrix))
	for i := range answer {
		answer[i] = make([]int, len(matrix[0]))
	}

	for i, v := range answer {
		for j := range v {
			if matrix[i][j] == -1 {
				answer[i][j] = lieMax[j]
			} else {
				answer[i][j] = matrix[i][j]
			}
		}
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	matrik := [][]int{
		{-1, 3, 4},
		{7, -1, 6},
		{9, -1, 2},
	}

	fmt.Println(modifiedMatrix(matrik))
}
