package main

func maxProfit(prices []int) int {
	n := len(prices)

	if n == 1 {
		return 0
	}

	bp := make([][]int, n)
	for i := range bp {
		bp[i] = make([]int, 2)
	}

	bp[0][0] = 0
	bp[0][1] = -prices[0]

	bp[1][0] = max(bp[0][0], bp[0][1]+prices[1])
	bp[1][1] = max(bp[0][1], bp[0][0]-prices[1])

	for i := 2; i < n; i++ {
		bp[i][0] = max(max(bp[i-2][1]+prices[i-2], bp[i-1][1]+prices[i-1]), bp[i-1][0])
		bp[i][1] = max(bp[i-1][0]-prices[i-1], bp[i-1][1])
	}

	return bp[n-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
