package main

import (
	"fmt"
)

func main() {
	words := []string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"}
	fmt.Println(maxProduct(words))
}

func maxProduct(words []string) int {
	wordMap := make(map[string]int)
	for _, word := range words {
		val := 0
		for _, ch := range word {
			val |= 1 << (ch - 'a')
		}
		wordMap[word] = val
	}

	answer := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if wordMap[words[i]]&wordMap[words[j]] == 0 {
				if x := len(words[i]) * len(words[j]); answer < x {
					answer = x
				}
			}
		}
	}
	return answer
}
