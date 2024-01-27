package main

import "fmt"

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	answer := []byte{}
	num := 0
	for n != 0 {
		answer = append(answer, byte(n%10)+'0')
		n = n / 10
		num++
		if num == 3 {
			answer = append(answer, '.')
			num = 0
		}

	}

	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}

	if answer[0] == '.' {
		answer = answer[1:]
	}

	return string(answer)
}

func main() {
	n := 123456789
	fmt.Println(thousandSeparator(n))
}
