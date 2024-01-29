package main

import (
	"fmt"
	"strings"
)

func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	answer := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			continue
		} else {
			answer++
		}
	}
	return answer
}

func main() {
	s := "bijdbisd"
	fmt.Println(countKeyChanges(s))
}
