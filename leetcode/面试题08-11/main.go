package main

import "fmt"

func waysToChange(n int) int {
	nums := []int{1, 5, 10, 25}

	mod := int(1e9 + 7)
	dp := make([]int, n+1)
	dp[0] = 1
	for _, coin := range nums {
		for j := coin; j <= n; j++ {
			dp[j] = (dp[j] + dp[j-coin]) % mod
		}
	}
	return dp[n]
}

func main() {
	n := 18
	fmt.Println(waysToChange(n))
}
