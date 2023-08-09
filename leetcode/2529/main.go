package main

import "fmt"

func maximumCount(nums []int) int {
	plus := 0
	minus := 0
	for _, v := range nums {
		if v > 0 {
			plus++
		}

		if v < 0 {
			minus++
		}
	}

	if plus > minus {
		return plus
	}

	return minus
}

func main() {
	nums := []int{1, 2, 3, -4, -5, 9, -9, 0, 0, -3}
	fmt.Println(maximumCount(nums))
}
