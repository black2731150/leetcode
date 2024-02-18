package main

import "fmt"

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		if i+reverse(i) == num {
			return true
		}
	}

	return false
}

func reverse(x int) int {
	answer := 0
	for x > 0 {
		answer = answer*10 + x%10
		x = x / 10
	}
	return answer
}

func main() {
	num := 443
	fmt.Println(sumOfNumberAndReverse(num))
}
