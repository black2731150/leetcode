package main

import "fmt"

func addBinary(a string, b string) string {
	var max int
	if len(a) > len(b) {
		max = len(a) - 1
	} else {
		max = len(b) - 1
	}

	alast := len(a) - 1
	blast := len(b) - 1

	var x, y byte
	var answer []byte
	var count byte = '0'

	for alast >= 0 || blast >= 0 {

		if alast >= 0 {
			x = a[alast]
		} else {
			x = '0'
		}

		if blast >= 0 {
			y = b[blast]
		} else {
			y = '0'
		}

		z := x + y + count - '0' - '0'
		count = '0'

		if z == '3' {
			z = '1'
			count = '1'
		} else if z == '2' {
			z = '0'
			count = '1'
		}

		answer = append(answer, z)
		max--

		alast--
		blast--
	}

	if count > '0' {
		answer = append(answer, '1')
	}

	retuenAnswer := string(answer)
	retuenAnswer = Reverse(retuenAnswer)
	return retuenAnswer
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	a := "1010101"
	b := "101011110110"

	fmt.Println(addBinary(a, b))
}
