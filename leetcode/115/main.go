package main

import "fmt"

func numDistinct(s string, t string) int {
	lens := len(s)
	lent := len(t)

	dp := make([][]int, lens+1)
	for i := range dp {
		dp[i] = make([]int, lent+1)
		dp[i][lent] = 1
	}

	for i := lens - 1; i >= 0; i-- {
		for j := lent - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j] + dp[i+1][j+1]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

func main() {
	s := "dihsdas"
	j := "shdjs"
	fmt.Println(numDistinct(s, j))
}
