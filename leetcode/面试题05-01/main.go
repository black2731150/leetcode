package main

import "fmt"

func insertBits(N int, M int, i int, j int) int {
	answer := 0

	for k := 0; k < i; k++ {
		answer = answer | (1&(N>>k))<<k
	}

	for k := i; k <= j; k++ {
		answer = answer | (1&(M>>(k-i)))<<k
	}

	for k := j + 1; k < 31; k++ {
		answer = answer | (1&(N>>k))<<k
	}

	return answer
}

func main() {
	N := 1024
	M := 19
	i := 2
	j := 6
	fmt.Println(insertBits(M, N, i, j))
}
