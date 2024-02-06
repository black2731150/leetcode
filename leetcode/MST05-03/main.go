package main

import "fmt"

func reverseBits(num int) int {

	answer := 0
	for i := 0; i < 32; i++ {

		x := 0
		hasOne := false
		for j := i; j < 32; j++ {
			if 1&(num>>j) == 1 {
				x++
			} else {
				if !hasOne {
					x++
					hasOne = true
				} else {
					break
				}
			}
		}
		answer = max(answer, x)

	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	num := 32312
	fmt.Println(reverseBits(num))
}
