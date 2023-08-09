package main

func generate(numRows int) [][]int {
	var answer [][]int

	answer = append(answer, []int{1})
	if numRows == 1 {
		return answer
	}

	for i := 1; i < numRows; i++ {
		A := []int{0}
		A = append(A, answer[i-1]...)

		B := answer[i-1]
		B = append(B, 0)

		C := []int{}

		for x := 0; x < len(A) && x < len(B); x++ {
			C = append(C, A[x]+B[x])
		}

		answer = append(answer, C)

	}

	return answer
}
