package main

import "fmt"

func exchangeBits(num int) int {
	answer := 0
	for i := 0; i < 32; i += 2 {
		x := 1 & (num >> i)
		y := 1 & (num >> (i + 1))

		answer = answer | (x&1)<<(i+1)
		answer = answer | (y&1)<<(i)
	}
	return answer
}

func main() {
	num := 3
	fmt.Println(exchangeBits(num))
}
