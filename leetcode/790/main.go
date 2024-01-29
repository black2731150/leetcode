package main

import "fmt"

func numTilings(n int) int {
	mod := int(1e9 + 7)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if n == 1 {
		return dp[1]
	}
	dp[2] = 2
	if n == 2 {
		return dp[2]
	}
	dp[3] = 5
	if n == 3 {
		return dp[3]
	}

	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % mod
	}
	return dp[n]
}

func main() {
	n := 6
	fmt.Println(numTilings(n))
}
