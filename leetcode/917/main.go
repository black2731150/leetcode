package main

import "fmt"

func reverseOnlyLetters(s string) string {
	n := len(s)
	left := 0
	right := n - 1

	answer := []byte(s)

	for left < n {
		if isLetter(s[left]) {
			break
		} else {
			left++
		}
	}

	for right >= 0 {
		if isLetter(s[right]) {
			break
		} else {
			right--
		}
	}

	if left >= right {
		return string(answer)
	}

	for left < right && left < n && right < n {
		if isLetter(s[left]) && isLetter(s[right]) {
			answer[left] = s[right]
			answer[right] = s[left]
			left++
			right--
		} else if !isLetter(s[left]) {
			left++
		} else if !isLetter(s[right]) {
			right--
		}
	}

	return string(answer)

}

func isLetter(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

func main() {
	s := "ab-cd"
	fmt.Println(reverseOnlyLetters(s))
}
