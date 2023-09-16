package main

import "fmt"

func isHappy(n int) bool {
	theMap := make(map[int]bool)

	for n != 1 && !theMap[n] {
		theMap[n] = true
		n = xxx(n)
	}

	return n == 1
}

func xxx(n int) int {
	answer := 0
	for n > 0 {
		answer = answer + (n%10)*(n%10)
		n = n / 10
	}

	return answer
}

func main() {
	n := 10
	fmt.Println(isHappy(n))
}
