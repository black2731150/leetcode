package main

import "fmt"

func vowelStrings(words []string, left int, right int) int {
	answer := 0

	for i := left; i <= right; i++ {
		word := words[i]
		first := word[0]
		end := word[len(word)-1]
		if isYuan(first) && isYuan(end) {
			answer++
		}
	}
	return answer
}

func isYuan(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}

	return false
}

func main() {
	word := []string{"are", "amy", "u"}
	fmt.Println(vowelStrings(word, 0, 2))
}
