package main

func shortestToChar(s string, c byte) []int {
	left, right := -1, -1
	for i := range s {
		if s[i] == c {
			right = i
			break
		}
	}

	answer := make([]int, len(s))
	for i := range answer {
		if s[i]==c{
			left=right
			right=
		}

		if i < left {
			answer[i] = left - i
			continue
		}


	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
