package main

import (
	"fmt"
	"strings"
)

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")

	if len(words) == 1 {
		if words[0][0] == words[0][len(words[0])-1] {
			return true
		} else {
			return false
		}
	}

	first := words[0][0]
	last := words[0][len(words[0])-1]
	for i := 1; i < len(words); i++ {
		if words[i][0] == last {
			last = words[i][len(words[i])-1]
			continue
		} else {
			return false
		}
	}

	return last == first

}

func main() {
	sentence := "leetcode"
	fmt.Println(isCircularSentence(sentence))
}
