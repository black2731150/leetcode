package main

import "fmt"

func numSquares(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	bp := make([]int, n+1)
	bp[0] = 0

	for i := 1; i < len(bp); i++ {
		bp[i] = n + 1
		for _, v := range nums {
			if i-v >= 0 {
				bp[i] = min(bp[i], bp[i-v]+1)
			}
		}
	}
	return bp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n := 12
	fmt.Println(numSquares(n))
}
