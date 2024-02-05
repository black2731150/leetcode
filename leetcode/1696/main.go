package main

import (
	"fmt"
	"math"
)

func maxResult(nums []int, k int) int {
	n := len(nums)

	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MinInt32
	}
	dp[0] = nums[0]

	quque := []int{0}
	for i := 1; i < n; i++ {
		if quque[0] < i-k {
			quque = quque[1:]
		}
		dp[i] = dp[quque[0]] + nums[i]

		for len(quque) > 0 && dp[i] >= dp[quque[len(quque)-1]] {
			quque = quque[:len(quque)-1]
		}
		quque = append(quque, i)
	}
	return dp[n-1]
}

func main() {
	nums := []int{1, -1, -2, 4, -7, 3}
	k := 3
	fmt.Println(maxResult(nums, k))
}
