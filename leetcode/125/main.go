package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	//转换成小写字母
	s = strings.ToLower(s)
	var left, right int
	left = 0
	right = len(s) - 1

	for left < right {
		if !isLowLetter(s[left]) {
			left++
			continue
		}

		if !isLowLetter(s[right]) {
			right--
			continue
		}

		if s[left] == s[right] {
			left++
			right--
			continue
		} else {
			return false
		}
	}

	return true
}

func isLowLetter(letter byte) bool {
	if (letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9') {
		return true
	}
	return false
}
func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
