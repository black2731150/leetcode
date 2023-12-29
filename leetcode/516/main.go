package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
		mat[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				if l == 2 {
					mat[i][j] = 2
				} else {
					mat[i][j] = mat[i+1][j-1] + 2
				}
			} else {
				mat[i][j] = max(mat[i+1][j], mat[i][j-1])
			}
		}
	}
	return mat[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}
