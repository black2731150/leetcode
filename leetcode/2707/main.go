package main

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	theMap := make(map[string]struct{})
	for _, v := range dictionary {
		theMap[v] = struct{}{}
	}

	dp := make([]int, len(s)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = len(s)
	}

	for i := 1; i < len(s); i++ {
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if _, find := theMap[s[j:i]]; find {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}

	return dp[len(s)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s := "leetscode"
	dictionary := []string{"leet", "code", "leetcode"}
	fmt.Println(minExtraChar(s, dictionary))
}
