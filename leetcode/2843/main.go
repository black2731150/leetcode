package main

import "fmt"

func countSymmetricIntegers(low int, high int) int {
	answer := 0
	for i := low; i <= high; i++ {
		if is(i) {
			answer++
		}
	}

	return answer

}

func is(x int) bool {
	cx := x
	digit := 0
	for cx > 0 {
		cx = cx / 10
		digit++
	}

	if digit%2 != 0 {
		return false
	}

	n := digit / 2
	right := 0
	for i := 0; i < n; i++ {
		right = right + x%10
		x = x / 10
	}

	left := 0
	for i := 0; i < n; i++ {
		left = left + x%10
		x = x / 10
	}

	return left == right
}

func main() {
	low := 1
	highL := 100
	fmt.Println(countSymmetricIntegers(low, highL))
}
