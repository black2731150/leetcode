package main

import "fmt"

func findDerangement(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	bp := make([]int, n+1)
	mod := int(1e9 + 7)

	bp[0] = 0
	bp[1] = 0
	bp[2] = 1

	for i := 3; i <= n; i++ {
		bp[i] = (i - 1) * ((bp[i-1]) + bp[i-2]) % mod
	}

	return bp[n]
}

func main() {
	n := 7
	fmt.Println(findDerangement(n))
}
