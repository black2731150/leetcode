package main

import "fmt"

func splitWordsBySeparator(words []string, separator byte) []string {
	answer := []string{}
	for _, word := range words {

		start, end := 0, 0
		for end < len(word) {
			if word[end] == separator {
				if start != end {
					answer = append(answer, word[start:end])
				}
				end++
				start = end
			} else {
				end++
			}
		}
		if start != end {
			answer = append(answer, word[start:end])
		}
	}

	return answer
}

func main() {
	sords := []string{"one.two.three", "four.five", "six"}
	separator := byte('.')
	fmt.Println(splitWordsBySeparator(sords, separator))
}
