package main

import "fmt"

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}

	if n%2 == 1 {
		return n
	} else {
		return n / 2
	}
}

func main() {
	n := 10
	fmt.Println(numberOfCuts(n))
}
