package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	bp := make([]int, len(nums))
	bp[0] = nums[0]
	bp[1] = max(nums[0], nums[1])
	bp[2] = max(bp[0]+nums[2], nums[1])

	for i := 3; i < len(nums); i++ {
		bp[i] = max(bp[i-2]+nums[i], bp[i-1])
	}
	return bp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
