package main

import "fmt"

func hanota(A []int, B []int, C []int) []int {
	nA := len(A)
	nB := len(B)
	if nA == 1 {
		C = append(C, A[0])
		return C
	}

	B = hanota(A[1:], C, B)
	C = append(C, A[0])
	A = A[:0]
	C = hanota(B[nB:nB+nA-1], A, C)

	return C
}

func main() {
	A := []int{13, 2, 1}
	B := []int{}
	C := []int{}
	fmt.Println(hanota(A, B, C))
}
