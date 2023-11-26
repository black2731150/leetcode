package main

import "fmt"

func sumOfMultiples(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum = sum + i
			continue
		}
	}

	return sum
}

func main() {
	n := 10
	fmt.Println(sumOfMultiples(n))
}
