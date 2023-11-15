package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for i := 0; i <= 32; i++ {
		if n%2 == 0 || n == 0 || n == 1 {
			n = n / 2
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	n := 1725
	fmt.Println(isPowerOfTwo(n))
}
