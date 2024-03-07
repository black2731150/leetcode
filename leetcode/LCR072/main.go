package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 0, x/2

	for left <= right {
		mid := (left + right) / 2

		m := mid * mid

		if m <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

func main() {
	x := 6
	fmt.Println(mySqrt(x))
}
