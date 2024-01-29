package main

import "fmt"

func minCost(costs [][]int) int {
	dp := make([][3]int, len(costs))

	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]

	for i := 1; i < len(costs); i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}

	return min(dp[len(dp)-1][0], min(dp[len(dp)-1][1], dp[len(dp)-1][2]))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	fmt.Println(minCost(costs))
}
