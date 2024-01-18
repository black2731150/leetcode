package main

import "fmt"

func maxProfit(prices []int) int {
	bp := make([][]int, len(prices))
	for i := range bp {
		bp[i] = make([]int, 2)
	}

	bp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		bp[i][0] = max(bp[i-1][0], bp[i-1][1]+prices[i])
		bp[i][1] = max(bp[i-1][1], -prices[i]+bp[i-1][0])
	}

	return bp[len(bp)-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
