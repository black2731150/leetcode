package main

import "fmt"

func countPrefixes(words []string, s string) int {
	answer := 0
	for _, word := range words {
		suucess := true

		if len(word) > len(s) {
			continue
		}

		for j := 0; j < len(word); j++ {
			if s[j] == word[j] {
				continue
			} else {
				suucess = false
				break
			}
		}

		if suucess {
			answer++
		}
	}
	return answer
}

func main() {
	words := []string{"a", "b", "c", "ab", "bc", "abc"}
	s := "abc"
	fmt.Println(countPrefixes(words, s))
}
