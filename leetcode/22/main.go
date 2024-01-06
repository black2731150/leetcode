package main

import "fmt"

func generateParenthesis(n int) []string {
	answer := []string{}

	var help func(left string, right string, get string)
	help = func(left string, right string, get string) {
		if len(left) == 0 && len(right) == 0 {
			answer = append(answer, get)
			return
		}

		if len(left) == 0 {
			help(left, right[1:], get+")")
			return
		}

		if len(left) == len(right) {
			help(left[:len(left)-1], right, get+"(")
		} else if len(left) <= len(right) {
			help(left, right[1:], get+")")
			help(left[:len(left)-1], right, get+"(")
		}
	}

	left, right := "", ""
	for i := 0; i < n; i++ {
		left += "("
		right += ")"
	}

	help(left, right, "")

	return answer
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
