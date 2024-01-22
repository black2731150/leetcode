package main

import "fmt"

func minAddToMakeValid(s string) int {
	stack := []byte{}

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, '(')
		} else {
			if len(stack) > 0 {
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, ')')
				}
			} else {
				stack = append(stack, ')')
			}
		}
	}

	return len(stack)
}

func main() {
	s := "())"
	fmt.Println(minAddToMakeValid(s))
}
