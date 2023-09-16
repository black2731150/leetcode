package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	numStack := []int{}
	for i := range tokens {
		if isSuan(tokens[i]) {
			num1 := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			num2 := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			if tokens[i] == "+" {
				numStack = append(numStack, num2+num1)
				continue
			}

			if tokens[i] == "-" {
				numStack = append(numStack, num2-num1)
				continue
			}

			if tokens[i] == "*" {
				numStack = append(numStack, num2*num1)
				continue
			}

			if tokens[i] == "/" {
				numStack = append(numStack, num2/num1)
				continue
			}
		} else {
			num, _ := strconv.Atoi(tokens[i])
			numStack = append(numStack, num)
		}
	}

	return numStack[0]
}

func isSuan(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}
