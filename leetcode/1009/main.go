package main

import "fmt"

func bitwiseComplement(n int) int {

	if n == 0 {
		return 1
	}

	answer := n
	for i := 0; i < 32 && n > 0; i++ {
		answer = answer ^ (1 << i)
		n = n / 2
	}

	return answer
}

func main() {
	n := 11
	fmt.Println(bitwiseComplement(n))
}
