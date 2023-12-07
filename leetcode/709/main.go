package main

import (
	"fmt"
	"unicode"
)

func toLowerCase(s string) string {
	answer := ""
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			answer += string(unicode.ToLower(rune(s[i])))
		} else {
			answer += string(s[i])
		}
	}
	return answer
}

func main() {
	s := "Hello"
	fmt.Println(toLowerCase(s))
}
