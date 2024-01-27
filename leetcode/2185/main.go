package main

import "fmt"

func prefixCount(words []string, pref string) int {
	answer := 0
	for _, word := range words {
		if len(word) >= len(pref) {
			if word[:len(pref)] == pref {
				answer++
			}
		}
	}
	return answer
}

func main() {
	words := []string{"pay", "attention", "practice", "attend"}
	perf := "at"
	fmt.Println(prefixCount(words, perf))
}
