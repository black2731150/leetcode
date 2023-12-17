package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	nums := make([]int, n*n)

	for _, v := range grid {
		for _, v2 := range v {
			nums[v2-1]++
		}
	}

	answer := []int{-1, -1}
	for i, v := range nums {
		if v == 2 {
			answer[0] = i + 1
		}

		if v == 0 {
			answer[1] = i + 1
		}
	}

	return answer
}
