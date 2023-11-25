package main

import (
	"fmt"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	n := len(words)
	answer := make([]string, n)
	for _, v := range words {
		lastChar := v[len(v)-1]
		num := int(lastChar - '0')
		answer[num-1] = v[:len(v)-1]
	}

	return strings.Join(answer, " ")
}

func main() {
	s := "is2 sentence4 This1 a3"
	fmt.Println(sortSentence(s))
}
