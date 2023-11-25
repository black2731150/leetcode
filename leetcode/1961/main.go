package main

import "fmt"

func isPrefixString(s string, words []string) bool {
	answer := ""
	for _, v := range words {
		answer = answer + v
		if answer == s {
			return true
		}
	}
	return false
}

func main() {
	s := "iloveyou"
	words := []string{"i", "love", "words"}
	fmt.Println(isPrefixString(s, words))
}
