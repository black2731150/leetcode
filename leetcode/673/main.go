package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	// dp 用于计算最长递增子序列
	dp := make([]int, n)
	// cnt 用于计算最长最长递增子序列的长度
	cnt := make([]int, n)

	dp[0] = 1

	maxLen := 0

	answer := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] = cnt[i] + cnt[j]
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
			answer = cnt[i]
		} else if dp[i] == maxLen {
			answer = answer + cnt[i]
		}
	}
	return answer
}

func main() {
	nums := []int{1, 2, 4, 3, 5, 4, 7, 2}
	fmt.Println(findNumberOfLIS(nums))
}
