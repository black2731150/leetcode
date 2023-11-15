package main

import "fmt"

func evenOddBit(n int) []int {
	even := 0
	odd := 0
	for i := 0; i < 10; i++ {
		x := 1 << i
		if n|x == n {
			if i%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	return []int{even, odd}
}

func main() {
	n := 7
	fmt.Println(evenOddBit(n))
}
