package main

import "fmt"

func addMinimum(word string) int {
	bp := make([]int, len(word))

	bp[0] = 2

	for i := 1; i < len(word); i++ {
		if word[i-1] < word[i] {
			bp[i] = bp[i-1] - 1
		} else {
			bp[i] = bp[i-1] + 2
		}
	}

	return bp[len(word)-1]
}

func main() {
	word := "abbbcab"
	fmt.Println(addMinimum(word))
}
