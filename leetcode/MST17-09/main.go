package main

import "fmt"

func getKthMagicNumber(k int) int {
	dp := []int{}
	dp = append(dp, 1)
	dp = append(dp, 3)
	dp = append(dp, 5)
	dp = append(dp, 7)
	p3, p5, p7 := 1, 1, 1
	for i := 4; i < k; i++ {
		dp = append(dp, 0)
		dp[i] = min(min(dp[p3]*3, dp[p5]*5), dp[p7]*7)
		if 3*dp[p3] == dp[i] {
			p3++
		}
		if 5*dp[p5] == dp[i] {
			p5++
		}
		if 7*dp[p7] == dp[i] {
			p7++
		}
	}

	return dp[k-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	k := 251
	fmt.Println(getKthMagicNumber(k))
}
