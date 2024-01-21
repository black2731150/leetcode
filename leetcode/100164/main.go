package main

func minimumArrayLength(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	minNum := 0
	for _, v := range nums {
		if v == min {
			minNum++
		}
	}

	if minNum%2 == 1 {
		return minNum/2 + 1
	} else {
		return minNum / 2
	}
}
