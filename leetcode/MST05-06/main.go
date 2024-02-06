package main

import "fmt"

func convertInteger(A int, B int) int {
	answer := 0
	for i := 0; i < 32; i++ {
		if 1&(A>>i) == 1&(B>>i) {
			continue
		} else {
			answer++
		}
	}
	return answer
}

func main() {
	A := 29
	B := 15
	fmt.Println(convertInteger(A, B))
}
