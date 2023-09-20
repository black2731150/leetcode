package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var answer uint32 = 0

	for i := 0; i < 32; i++ {
		if num&(1<<i) == 1<<i {
			answer = answer | (1 << (31 - i))
		}
	}
	return answer
}

func main() {
	var num uint32 = 123

	fmt.Println(reverseBits(num))
}
