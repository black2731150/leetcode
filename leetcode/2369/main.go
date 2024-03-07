package main

import "fmt"

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, len(nums))

	dp[0] = false
	dp[1] = check(nums[:2])

	if n >= 3 {
		dp[2] = check(nums[:3])
	}

	for i := 3; i < n; i++ {
		a := dp[i-2] && check(nums[i-1:i+1])
		b := dp[i-3] && check(nums[i-2:i+1])
		dp[i] = a || b
	}

	return dp[n-1]
}

func check(nums []int) bool {
	n := len(nums)
	if n == 2 {
		return nums[0] == nums[1]
	}

	if n == 3 {
		a := nums[0] == nums[1] && nums[1] == nums[2]
		b := nums[0]+1 == nums[1] && nums[1]+1 == nums[2]
		return a || b
	}

	return false
}

func main() {
	nums := []int{4, 4, 4, 5, 6}
	fmt.Println(validPartition(nums))
}
