package main

func findMagicIndex(nums []int) int {
	for i, v := range nums {
		if v == i {
			return i
		}
	}

	return -1
}
