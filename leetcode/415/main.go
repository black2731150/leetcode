package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}

	if len(num2) == 0 {
		return num1
	}

	var answer []byte

	num1Last := len(num1) - 1
	num2Last := len(num2) - 1

	var n1, n2 byte = '0', '0'
	var count byte = '0'
	for num1Last >= 0 || num2Last >= 0 || count != '0' {

		if num1Last < 0 {
			n1 = '0'
		} else {
			n1 = num1[num1Last]
		}

		if num2Last < 0 {
			n2 = '0'
		} else {
			n2 = num2[num2Last]
		}

		c := n1 + n2 + count - '0' - '0'
		count = '0'

		if c > '9' {
			c = c - 10
			count = '1'
		}

		answer = append(answer, c)
		num1Last--
		num2Last--
	}

	reverseString(answer)

	return string(answer)
}

func reverseString(s []byte) {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		temp := s[i]
		s[i] = s[sLen-1-i]
		s[sLen-1-i] = temp
	}
}

func main() {
	num1 := "11"
	num2 := "12"
	fmt.Println(addStrings(num1, num2))
}
