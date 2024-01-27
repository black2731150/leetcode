package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	word1Map, word2Map := make(map[string]int), make(map[string]int)
	for i := range words1 {
		word1Map[words1[i]]++
	}

	for i := range words2 {
		word2Map[words2[i]]++
	}

	answer := 0
	for k1, num1 := range word1Map {
		if num1 == 1 {
			if num2, find := word2Map[k1]; find {
				if num1 == num2 {
					answer++
				}
			}
		}
	}

	return answer
}

func main() {
	words1 := []string{"a", "ab"}
	words2 := []string{"a", "a", "a", "ab"}
	fmt.Println(countWords(words1, words2))
}
