package main

import "fmt"

func freqAlphabets(s string) string {
	answer := []byte{}

	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			n1 := answer[len(answer)-1]
			n2 := answer[len(answer)-2]
			answer = answer[:len(answer)-2]

			x := int(n1-'a') + int(n2-'a')*10
			answer = append(answer, byte(x+'a'))
		} else {
			answer = append(answer, s[i]-'0'+'a')
		}
	}
	for i := range answer {
		answer[i]--
	}

	return string(answer)
}

func main() {
	s := "10#11#12"
	fmt.Println(freqAlphabets(s))
}
