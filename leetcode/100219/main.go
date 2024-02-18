package main

import "fmt"

func maxPalindromesAfterOperations(words []string) int {
	theMap := make(map[byte]int)
	for _, word := range words {
		for j := 0; j < len(word); j++ {
			theMap[word[j]]++
		}
	}

	answer := 0
	for _, word := range words {
		cm := make(map[byte]int, len(theMap))
		for k, v := range theMap {
			cm[k] = v
		}
		n := len(word)
		if n%2 == 1 {
			for k, v := range cm {
				if v%2 == 1 {
					cm[k]--
					n--
					break
				}
			}

			if n%2 == 1 {
				for k, v := range cm {
					if v > 0 {
						cm[k]--
						n--
						break
					}
				}
			}

			if n%2 == 1 {
				continue
			}
		}

		if n == 0 {
			answer++
			theMap = cm
			continue
		}

		if n%2 == 0 {
			for k, v := range cm {
				if v >= 2 {
					for cm[k] >= 2 && n > 0 {
						cm[k] -= 2
						n -= 2
					}
				}
			}
		}

		if n == 0 {
			answer++
			theMap = cm
		}
	}

	return answer
}

func main() {
	words := []string{"acc", "bc"}
	fmt.Println(maxPalindromesAfterOperations(words))
}
