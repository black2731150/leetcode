package main

import (
	"fmt"
	"math"
)

func pivotInteger(n int) int {
	t := (n*n + n) / 2
	x := int(math.Sqrt(float64(t)))
	if x*x == t {
		return x
	}
	return -1
}

func main() {
	n := 8
	fmt.Println(pivotInteger(n))
}
