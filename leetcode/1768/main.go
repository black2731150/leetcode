package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	answer := []byte{}
	for i < len(word1) && i < len(word2) {
		answer = append(answer, word1[i])
		answer = append(answer, word2[i])
		i++
	}

	if i == len(word1) {
		answer = append(answer, []byte(word2[i:])...)
	} else {
		answer = append(answer, []byte(word1[i:])...)
	}
	return string(answer)
}

func main() {
	word1 := "abc"
	word2 := "jhg"
	fmt.Println(mergeAlternately(word1, word2))
}
