package main

import "fmt"

func areOccurrencesEqual(s string) bool {
	letters := make(map[byte]int)
	for i := range s {
		letters[s[i]]++
	}

	x := letters[s[0]]
	for _, v := range letters {
		if x != v {
			return false
		}
	}

	return true
}

func main() {
	s := "abacbc"
	fmt.Println(areOccurrencesEqual(s))
}
