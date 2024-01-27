package main

import "fmt"

func removeDuplicateLetters(s string) string {
	bytes := make([]int, 26)
	for i := range s {
		bytes[s[i]-'a']++
	}

	have := make([]bool, 26)
	answer := []byte{}
	for i := range s {
		if len(answer) == 0 {
			answer = append(answer, s[i])
			have[s[i]-'a'] = true
		} else {
			if s[i] > answer[len(answer)-1] && !have[s[i]-'a'] {
				answer = append(answer, s[i])
				have[s[i]-'a'] = true
			} else {
				for len(answer) > 0 && s[i] <= answer[len(answer)-1] && bytes[answer[len(answer)-1]-'a'] > 0 && !have[s[i]-'a'] {
					last := answer[len(answer)-1]
					answer = answer[:len(answer)-1]
					have[last-'a'] = false
				}

				if !have[s[i]-'a'] {
					answer = append(answer, s[i])
					have[s[i]-'a'] = true
				}
			}
		}
		bytes[s[i]-'a']--
	}

	return string(answer)
}

func main() {
	s := "cbacdcbc"
	fmt.Println(removeDuplicateLetters(s))
}
