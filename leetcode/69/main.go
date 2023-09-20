package main

import "fmt"

func mySqrt(x int) int {
	i := 0
	for ; i*i < x; i++ {
	}

	if i*i == x {
		return i
	}
	return i - 1
}

func main() {
	x := 7
	fmt.Println(mySqrt(x))
}
