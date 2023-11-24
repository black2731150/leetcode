package main

import "fmt"

func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}

	maxLen := 0
	left, repeat := 0, 0

	for right := 1; right < n; right++ {
		if s[right] == s[right-1] {
			repeat++
		}

		if repeat == 2 {
			left++
			for s[left] != s[left-1] {
				left++
			}
			repeat = 1
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "1111111111"
	fmt.Println(longestSemiRepetitiveSubstring(s))
}
