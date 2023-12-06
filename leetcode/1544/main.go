package main

import (
	"fmt"
	"unicode"
)

func makeGood(s string) string {

	for i := 0; i < len(s)-1; {
		if s[i] != s[i+1] {
			if s[i] == byte(unicode.ToLower(rune(s[i+1]))) {
				s = s[:i] + s[i+2:]
				i = 0
				continue
			}

			if s[i] == byte(unicode.ToUpper(rune(s[i+1]))) {
				s = s[:i] + s[i+2:]
				i = 0
				continue
			}
			i++
		} else {
			i++
		}
	}

	return s
}

func main() {
	s := "abBAcC"
	fmt.Println(makeGood(s))
}
