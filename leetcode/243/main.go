package main

import "fmt"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	word1Index := []int{}
	for i, word := range wordsDict {
		if word == word1 {
			word1Index = append(word1Index, i)
		}
	}

	answer := len(wordsDict)

	for i, word := range wordsDict {
		if word == word2 {
			for _, word1i := range word1Index {
				answer = min(answer, abs(word1i-i))
			}
		}
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "coding"
	word2 := "practice"
	fmt.Println(shortestDistance(wordsDict, word1, word2))
}
