package main

import "fmt"

func firstPalindrome(words []string) string {
	for _, word := range words {
		n := len(word)
		if word[0] == word[n-1] {
			left := 0
			right := n - 1
			is := true
			for left <= right {
				if word[left] == word[right] {
					left++
					right--
				} else {
					is = false
					break
				}
			}
			if is {
				return word
			}
		}
	}

	return ""
}

func main() {
	s := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(s))
}
