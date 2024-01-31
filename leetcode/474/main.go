package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for k := range strs {
		z, o := zeroAndOne(strs[k])
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = max(dp[i-z][j-o]+1, dp[i][j])
			}
		}
	}
	return dp[m][n]
}

func zeroAndOne(s string) (zero, one int) {
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	fmt.Println(findMaxForm(strs, m, n))
}
