package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	n := len(s)
	left := 0
	right := n - 1
	answer := make([]byte, len(s))
	for left <= right {
		if s[left] != s[right] {
			if s[left] < s[right] {
				answer[left] = byte(s[left])
				answer[right] = byte(s[left])
			} else {
				answer[left] = byte(s[right])
				answer[right] = byte(s[right])
			}
		} else {
			answer[left] = byte(s[left])
			answer[right] = byte(s[right])
		}

		left++
		right--
	}

	return string(answer)
}

func main() {
	s := "egcfe"
	fmt.Println(makeSmallestPalindrome(s))
}
