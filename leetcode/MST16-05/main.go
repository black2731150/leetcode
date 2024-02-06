package main

import "fmt"

func trailingZeroes(n int) int {
	x := 5
	answer := 0
	for n/x > 0 {
		answer += n / x
		x = x * 5
	}
	return answer
}

func main() {
	n := 13
	fmt.Println(trailingZeroes(n))
}
