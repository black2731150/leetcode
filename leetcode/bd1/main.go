package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var input string
		fmt.Scanln(&input)
		answer := f(input)
		fmt.Println(answer)
	}
}

func f(s string) string {
	n := len(s)
	if n == 1 || n == 2 {
		return s
	}

	if n == 3 {
		if s[0] == s[1] && s[1] == s[2] {
			return s[:2]
		}
	}

	for i := 3; i < len(s); i++ {
		if s[i] == s[i-1] && s[i-1] == s[i-2] {
			s = s[:i] + s[i+1:]
			i--
			continue
		}

		if s[i] == s[i-1] && s[i-2] == s[i-3] {
			s = s[:i-1] + s[i:]
			i--
			continue
		}
	}
	return s
}
