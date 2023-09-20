package main

import "fmt"

func hammingWeight(num uint32) int {
	answer := 0

	for i := 0; i < 32; i++ {
		if num&(1<<i) == (1 << i) {
			answer++
		}
	}

	return answer
}

func main() {
	var num uint32 = 123
	fmt.Println(hammingWeight(num))
}
