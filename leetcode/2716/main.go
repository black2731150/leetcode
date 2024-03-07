package main

import "fmt"

func minimizedStringLength(s string) int {
	chars := 0

	for i := 0; i < len(s); i++ {
		chars = chars | (1 << int(s[i]-'a'))
	}

	answer := 0
	for i := 0; i < 26; i++ {
		if (chars>>i)&1 == 1 {
			answer++
		}
	}

	return answer
}

func main() {
	s := "ipi"
	fmt.Println(minimizedStringLength(s))
}
