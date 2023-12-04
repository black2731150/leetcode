package main

import "fmt"

func maxNonDecreasingLength(nums1 []int, nums2 []int) (ret int) {
	ret = 1
	n := len(nums1)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = 1
		dp[i][1] = 1
	}
	for i := 1; i < n; i++ {
		if nums1[i] >= nums1[i-1] {
			dp[i][0] = max(dp[i][0], dp[i-1][0]+1)
		}
		if nums1[i] >= nums2[i-1] {
			dp[i][0] = max(dp[i][0], dp[i-1][1]+1)
		}
		if nums2[i] >= nums1[i-1] {
			dp[i][1] = max(dp[i][1], dp[i-1][0]+1)
		}
		if nums2[i] >= nums2[i-1] {
			dp[i][1] = max(dp[i][1], dp[i-1][1]+1)
		}

		ret = max(ret, max(dp[i][0], dp[i][1]))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums1 := []int{2, 3, 1}
	nums2 := []int{1, 2, 1}
	fmt.Println(maxNonDecreasingLength(nums1, nums2))
}
