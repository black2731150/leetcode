package main

import "fmt"

func maximumCostSubstring(s string, chars string, vals []int) int {
	char2num := make(map[rune]int, 26)
	for i := 'a'; i <= 'z'; i++ {
		char2num[i] = int(i-'a') + 1
	}
	for i, char := range chars {
		char2num[char] = vals[i]
	}

	fmt.Println(char2num)

	answer := 0

	x := 0
	for _, char := range s {
		x = max(x, 0) + char2num[char]
		answer = max(answer, x)
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
	s := "aabc"
	chars := "ab"
	vals := []int{-1, -2}
	fmt.Println(maximumCostSubstring(s, chars, vals))
}
