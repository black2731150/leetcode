package main

import "fmt"

func countGoodStrings(low int, high int, zero int, one int) int {
	bp := make([]int, high+1)
	mod := int(1e9 + 7)

	if zero != one {
		bp[min(zero, one)] = 1
		if max(zero, one)%min(zero, one) == 0 {
			bp[max(zero, one)] = 2
		} else {
			bp[max(zero, one)] = 1
		}
	} else {
		bp[zero] = 2
	}

	for i := min(zero, one) + 1; i < max(zero, one); i++ {
		if i%min(zero, one) == 0 {
			bp[i] = 1
		}
	}

	for i := max(zero, one) + 1; i <= high; i++ {
		bp[i] = (bp[i-zero] + bp[i-one]) % mod
	}

	answer := 0
	for i := low; i <= high; i++ {
		answer = (answer + bp[i]) % mod
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	low := 3
	high := 5
	zero := 1
	one := 1
	fmt.Println(countGoodStrings(low, high, zero, one))
}
