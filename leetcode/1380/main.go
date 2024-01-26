package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	hangMins := make([]int, m)
	lieMaxs := make([]int, n)

	for i := 0; i < m; i++ {
		hangMins[i] = matrix[i][0]
	}

	for i := 0; i < n; i++ {
		lieMaxs[i] = matrix[0][i]
	}

	for i, v := range matrix {
		for j, val := range v {
			hangMins[i] = min(hangMins[i], val)
			lieMaxs[j] = max(lieMaxs[j], val)
		}
	}

	answer := []int{}
	for i, v := range matrix {
		for j, val := range v {
			if val == hangMins[i] && val == lieMaxs[j] {
				answer = append(answer, val)
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	matrix := [][]int{
		{3, 7, 8}, {9, 11, 13}, {15, 16, 17},
	}

	fmt.Println(luckyNumbers(matrix))
}
