package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	answer := []string{}

	for i := 1; i <= n; i++ {
		if isFizz(i) && isBuzz(i) {
			answer = append(answer, "FizzBuzz")
			continue
		} else {
			if isFizz(i) {
				answer = append(answer, "Fizz")
				continue
			}

			if isBuzz(i) {
				answer = append(answer, "Buzz")
				continue
			}

		}

		answer = append(answer, strconv.Itoa(i))
	}

	return answer
}

func isFizz(n int) bool {
	return n%3 == 0
}

func isBuzz(n int) bool {
	return n%5 == 0
}

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}
