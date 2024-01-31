package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	N := len(nums)

	dp := make([]int, target+1)
	dp[0] = 1

	for v := 1; v <= target; v++ {
		for i := 0; i < N; i++ {
			if v-nums[i] >= 0 {
				dp[v] += dp[v-nums[i]]
			}
		}
	}

	return dp[target]
}

func main() {
	nums := []int{3, 1, 2, 4}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}
