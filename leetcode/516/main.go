package main

func longestPalindromeSubseq(s string) int {
	max := 0

	for i := 0; i < len(s); i++ {

	}
}

func help(s string, left, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left++
			right--
		} else {
			break
		}
	}
	return right - left + 1
}
