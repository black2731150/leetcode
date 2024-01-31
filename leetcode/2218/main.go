package main

import "fmt"

// 分组背包问题
func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	weight := make([][]int, n)
	for i := range weight {
		weight[i] = make([]int, len(piles[i])+1)
	}
	for i := range piles {
		for j := range piles[i] {
			weight[i][j+1] = weight[i][j] + piles[i][j]
		}
	}

	dp := make([]int, k+1)

	for i := 0; i < n; i++ {
		for v := k; v >= 1; v-- {
			for j := 1; j <= v && j <= len(piles[i]); j++ {
				dp[v] = max(dp[v], dp[v-j]+weight[i][j])
			}
		}
	}

	return dp[k]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	piles := [][]int{{1, 100, 3}, {7, 8, 9}}
	k := 2
	fmt.Println(maxValueOfCoins(piles, k))
}
