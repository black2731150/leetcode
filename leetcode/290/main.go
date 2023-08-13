package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	p2w := map[string]string{}
	w2p := map[string]string{}

	for i := range pattern {
		_, find1 := p2w[string(pattern[i])]
		_, find2 := w2p[words[i]]

		if !find1 && find2 {
			p2w[string(pattern[i])] = words[i]
			w2p[words[i]] = string(pattern[i])
		} else {
			if p2w[string(pattern[i])] == words[i] && w2p[words[i]] == string(pattern[i]) {
				continue
			} else {
				return false
			}
		}

	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat fish"
	fmt.Println(wordPattern(pattern, s))
}
