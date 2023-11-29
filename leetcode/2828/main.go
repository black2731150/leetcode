package main

import "fmt"

func isAcronym(words []string, s string) bool {

	if len(words) != len(s) {
		return false
	}

	for i, v := range words {
		if s[i] != v[0] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"alice", "bob", "charlie"}
	s := "abc"
	fmt.Println(isAcronym(words, s))
}
