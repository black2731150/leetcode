package main

import "sort"

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)

	sum := 0
	n := len(arr)
	min := 10000

	answer := 0

	for i := range arr {
		theNum := (target - sum) / (n - i)

		j := target - (sum + theNum*(n-i))
		if j < 0 {
			break
		}
		if x := abs(j); x <= min && theNum >= 0 {
			answer = theNum
			min = x
		}
		sum = sum + theNum
	}

	return answer
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
