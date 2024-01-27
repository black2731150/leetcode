package main

import "fmt"

func removeVowels(s string) string {
	answer := []byte{}
	for i := range s {
		ch := s[i]
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			continue
		} else {
			answer = append(answer, ch)
		}
	}
	return string(answer)
}

func main() {
	s := "leetcodeisacommunityforcoders"
	fmt.Println(removeVowels(s))
}
