package main

import "fmt"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}

	answers := make([][]int, n)
	for i := range answers {
		answers[i] = make([]int, n)
	}

	copy(answers[0], matrix[0])

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			firstIndex, secondIndex, thredIndex := j-1, j, j+1
			if j-1 < 0 {
				firstIndex = 0
			}

			if j+1 >= n {
				thredIndex = n - 1
			}

			answers[i][j] = min(answers[i-1][firstIndex], answers[i-1][secondIndex], answers[i-1][thredIndex]) + matrix[i][j]
		}
	}

	ans := answers[len(answers)-1][0]
	for _, v := range answers[len(answers)-1] {
		if v < ans {
			ans = v
		}
	}

	return ans
}

func min(x, y, z int) int {
	if x < y {
		if z < x {
			return z
		} else {
			return x
		}
	} else {
		if z < y {
			return z
		} else {
			return y
		}
	}
}

func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(matrix))
}
