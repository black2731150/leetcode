package main

func lengthOfLIS(nums []int) int {
	bp := make([]int, len(nums))
	bp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			bp[i] = bp[i-1]
		} else {
			bp[i] = bp[i-1] + 1
		}
	}
	return bp[len(bp)-1]
}
