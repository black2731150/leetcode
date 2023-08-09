package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := range s {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	x := 123321
	fmt.Println(isPalindrome(x))
}
