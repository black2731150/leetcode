package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	n := len(s)

	if n == 1 {
		return 1
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "abaccc"
	fmt.Println(lengthOfLongestSubstringTwoDistinct(s))
}
