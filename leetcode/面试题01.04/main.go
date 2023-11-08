package main

import "fmt"

func canPermutePalindrome(s string) bool {
	themap := make(map[rune]int)
	for _, ch := range s {
		themap[ch]++
	}

	answer := 0
	for _, v := range themap {
		if v%2 == 0 {
			continue
		} else {
			answer++
		}
	}
	return answer <= 1
}

func main() {
	s := "ajdiadskbd"
	fmt.Println(canPermutePalindrome(s))
}
