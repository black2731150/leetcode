package main

import (
	"math"
	"strconv"
)

func superpalindromesInRange(left string, right string) int {
	l, _ := strconv.ParseInt(left, 10, 64)
	r, _ := strconv.ParseInt(right, 10, 64)
	start := int(math.Sqrt(float64(l)))
	end := int(math.Sqrt(float64(r)))

	answer := 0
	for i := start; i <= end; i++ {
		if isPalindrome(int64(i)) {
			q := int64(i) * int64(i)
			if isPalindrome(q) && q >= l && q <= r {
				answer++
			}
		}
	}
	return answer
}

func isPalindrome(x int64) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := int64(0)
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}
