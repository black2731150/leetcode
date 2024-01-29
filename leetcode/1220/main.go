package main

import "fmt"

func countVowelPermutation(n int) int {
	mod := int(1e9 + 7)

	dp := make([][5]int, n+1)
	dp[1][0] = 1
	dp[1][1] = 1
	dp[1][2] = 1
	dp[1][3] = 1
	dp[1][4] = 1

	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][1] % mod
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][0] + dp[i-1][1] + dp[i-1][3] + dp[i-1][4]) % mod
		dp[i][3] = (dp[i-1][2] + dp[i-1][4]) % mod
		dp[i][4] = dp[i-1][0] % mod
	}

	answer := 0
	for i := 0; i <= 4; i++ {
		answer = (answer + dp[n][i]) % mod
	}
	return answer
}

func main() {
	n := 5
	fmt.Println(countVowelPermutation(n))
}
