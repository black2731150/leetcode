package main

func scoreOfParentheses(s string) int {
	answer := 0
	stack := []byte{}

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, '(')
		} else {
			for j := i + 1; j < len(s); j++ {

			}
		}
	}
}
