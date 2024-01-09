package main

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}

	n := len(nums)

	lastUp := 0
	for i := 1; i < n; i++ {

		if nums[i] > nums[i-1] {
			lastUp = i - 1
			continue
		}
	}

	for i, j := 0, len(nums[lastUp:])-1; i < j; i, j = i+1, j-1 {
		nums[lastUp:][i], nums[lastUp:][j] = nums[lastUp:][j], nums[lastUp:][i]
	}

}
