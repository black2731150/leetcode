package main

import "fmt"

func numberOfSpecialSubstrings(s string) int {
	n := len(s)

	answer := 0

	for i := range s {
		theMap := make(map[byte]bool)
		for j := i; j < n; j++ {
			if _, find := theMap[s[j]]; find {
				break
			} else {
				theMap[s[j]] = true
				answer++
			}
		}
	}

	return answer

}

func main() {
	s := "abab"
	fmt.Println(numberOfSpecialSubstrings(s))
}
