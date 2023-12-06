package main

import "fmt"

func countLetters(s string) int {
	left, right := 0, 0

	answer := 0
	for right < len(s) {
		if s[left] == s[right] {
			right++
		} else {
			x := right - left
			answer += x * (x + 1) / 2
			left = right
		}
	}

	x := right - left
	answer += x * (x + 1) / 2

	return answer
}

func main() {
	s := "aaaba"
	fmt.Println(countLetters(s))
}
