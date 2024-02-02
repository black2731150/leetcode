package main

import "fmt"

func add(a int, b int) int {
	answer := int32(0)

	count := 0
	for i := 0; i < 32; i++ {
		x := 1 & (a >> i)
		y := 1 & (b >> i)

		if count == 0 {
			if x == 1 && y == 1 {
				count = 1
			} else if x == 0 && y == 0 {
				continue
			} else {
				answer = answer | (1 << i)
			}
		} else {
			if x == 1 && y == 1 {
				count = 1
				answer = answer | (1 << i)
			} else if x == 0 && y == 0 {
				count = 0
				answer = answer | (1 << i)
			} else {
				count = 1
			}
		}
	}

	return int(answer)
}

func main() {
	a := 1
	b := -2
	fmt.Println(add(a, b))
}
