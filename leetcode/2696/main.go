package main

import "fmt"

func minLength(s string) int {
	if len(s) == 1 {
		return 1
	}

	for i := 0; i < len(s)-1; i++ {
		if len(s) == 1 {
			return 1
		}

		if i < 0 {
			i = 0
		}
		if s[i] == 'A' && s[i+1] == 'B' {
			s = s[:i] + s[i+2:]
			i--
			i--
			continue
		}

		if s[i] == 'C' && s[i+1] == 'D' {
			s = s[:i] + s[i+2:]
			i--
			i--
			continue
		}
	}

	return len(s)
}

func main() {
	s := "AAAABBBBC"
	fmt.Println(minLength(s))
}
