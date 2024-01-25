package main

import "fmt"

func canPermutePalindrome(s string) bool {
	leeters := make([]int, 26)
	for i := range s {
		leeters[s[i]-'a']++
	}

	haveOne := 0
	for _, v := range leeters {
		if v%2 == 0 {
			continue
		} else {
			if haveOne == 1 {
				return false
			}
			haveOne++
		}
	}

	return true
}

func main() {
	s := "code"
	fmt.Println(canPermutePalindrome(s))
}
