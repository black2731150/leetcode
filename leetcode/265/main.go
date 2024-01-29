package main

import (
	"fmt"
	"math"
)

func minCostII(costs [][]int) int {
	n := len(costs)
	k := len(costs[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	for i := 0; i < k; i++ {
		dp[0][i] = costs[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			minNum := math.MaxInt
			for x := 0; x < k; x++ {
				if x != j {
					minNum = min(minNum, dp[i-1][x])
				}
			}

			dp[i][j] = minNum + costs[i][j]
		}
	}

	answer := math.MaxInt
	for i := 0; i < k; i++ {
		answer = min(answer, dp[n-1][i])
	}
	return answer
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	costs := [][]int{{1, 5, 3}, {2, 9, 4}}
	fmt.Println(minCostII(costs))
}
