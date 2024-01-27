package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(help(nums[1:]), help(nums[:len(nums)-1]))
}

func help(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}
	bp := make([]int, n)
	bp[0] = nums[0]
	bp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		bp[i] = max(bp[i-2]+nums[i], bp[i-1])
	}

	return bp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}
