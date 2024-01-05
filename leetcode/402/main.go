package main

import "fmt"

func removeKdigits(num string, k int) string {
	n := len(num)
	if n == k {
		return "0"
	}
	stack := []int{}

	for _, digit := range num {
		num := int(digit - '0')
		if len(stack) > 0 {
			for len(stack) > 0 && stack[len(stack)-1] != 0 && stack[len(stack)-1] > num && k > 0 {
				k--
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, num)
	}

	stack = stack[:len(stack)-k]

	answer := ""
	for len(stack) > 0 {
		if stack[0] != 0 {
			break
		} else {
			stack = stack[1:]
		}
	}
	for _, v := range stack {
		answer += string(rune(v + '0'))
	}

	if len(stack) == 0 {
		return "0"
	}
	return answer
}

func main() {
	num := "10200"
	k := 1
	fmt.Println(removeKdigits(num, k))
}
