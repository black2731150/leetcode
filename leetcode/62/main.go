package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	mat := make([][]int, m)
	for i := range mat {
		mat[i] = make([]int, n)
	}

	// 初始化右下角
	mat[m-1][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 初始化边界
			if i == m-1 || j == n-1 {
				mat[i][j] = 1
			} else {
				mat[i][j] = mat[i+1][j] + mat[i][j+1]
			}
		}
	}

	return mat[0][0]
}

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}
