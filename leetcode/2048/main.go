package main

import "fmt"

func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		if hep(i) {
			return i
		}
	}
}

func hep(x int) bool {
	digits := make([]int, 10)

	for x > 0 {
		digits[x%10]++
		x = x / 10
	}

	for i := range digits {
		if digits[i] > 0 && digits[i] != i {
			return false
		}
	}
	return true
}

func main() {
	n := 1
	fmt.Println(nextBeautifulNumber(n))
}
