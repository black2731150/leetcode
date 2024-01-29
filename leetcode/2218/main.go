package main

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	qianzhui := make([][]int, n)
	for i := range qianzhui {
		qianzhui[i] = make([]int, len(piles[i])+1)
	}

	for i := range piles {
		for j := range piles[i] {
			qianzhui[i][j+1] = qianzhui[i][j] + piles[i][j]
		}
	}

	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}

	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
