package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)

	n := len(beans)

	adds := make([]int64, n+1)
	for i := 1; i < len(beans)+1; i++ {
		adds[i] = adds[i-1] + int64(beans[i-1])
	}

	var minNum int64
	minNum = math.MaxInt64

	for i := 0; i < n; i++ {
		left := adds[i]
		right := adds[n] - adds[i+1] - int64(beans[i])*int64(n-i-1)
		minNum = min(minNum, left+right)
	}

	return minNum
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func main() {
	beans := []int{4, 1, 6, 5}
	fmt.Println(minimumRemoval(beans))
}
