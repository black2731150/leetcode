package main

import "fmt"

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	answer := 0

	for n > 0 {
		n = n / 5
		answer = answer + n
	}

	return answer
}

func main() {
	n := 20
	fmt.Println(trailingZeroes(n))
}
