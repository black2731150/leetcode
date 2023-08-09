package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var last = len(nums) - 1

	for i := 0; i <= last; i++ {
		if nums[i] == val {
			nums[i], nums[last] = nums[last], nums[i]
			last--
			i--
		}
	}

	return last + 1
}
