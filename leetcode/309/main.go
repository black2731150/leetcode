package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)

	if n == 1 {
		return 0
	}

	bp := make([][3]int, n)

	bp[0][0] = -prices[0]
	bp[0][1] = 0
	bp[0][2] = 0

	for i := 1; i < n; i++ {
		bp[i][0] = max(bp[i-1][0], bp[i-1][2]-prices[i])
		bp[i][1] = bp[i-1][0] + prices[i]
		bp[i][2] = max(bp[i-1][1], bp[i-1][2])
	}

	return max(bp[n-1][1], bp[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}
