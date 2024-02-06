package main

import (
	"fmt"
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	rA := math.MaxInt
	for i := 0; i < len(a); i++ {
		left, right := 0, len(b)-1
		for left <= right {
			mid := (left + right) / 2
			x := abs(a[i] - b[mid])
			rA = min(rA, x)

			if a[i] < b[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return rA
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	a := []int{1, 2, 11, 15}
	b := []int{4, 12, 19, 23, 127, 235}
	fmt.Println(smallestDifference(a, b))
}
