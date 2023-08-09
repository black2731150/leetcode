package main

func countBits(n int) []int {
	answer := []int{}
	for i := 0; i <= n; i++ {
		answer = append(answer, onesCount(i))
	}
	return answer
}

func onesCount(x int) int {
	onse := 0
	for ; x > 0; x = x & (x - 1) {
		onse++
	}
	return onse
}
