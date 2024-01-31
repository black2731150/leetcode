package main

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = max(dp[i-1][j-1]+1, dp[i][j])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}
