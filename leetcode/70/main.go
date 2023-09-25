package main

import "fmt"

func climbStairs(n int) int {
	nums := []int{0, 1, 2}
	if n <= 2 {
		return nums[n]
	}

	for i := 3; i <= n; i++ {
		nums = append(nums, nums[i-1]+nums[i-2])
	}

	return nums[len(nums)-1]

}

func main() {
	fmt.Println(climbStairs(7))
}
