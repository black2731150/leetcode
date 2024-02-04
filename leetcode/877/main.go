package main

import "fmt"

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(dp[i+1][j]+piles[i], dp[i][j-1]+piles[j])
		}
	}

	return dp[0][n-1] > 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	piles := []int{3, 2, 1, 3}
	fmt.Println(stoneGame(piles))
}
