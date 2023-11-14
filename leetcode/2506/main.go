package main

import "fmt"

func similarPairs(words []string) int {
	theMap := make(map[uint32]int)

	for _, v := range words {
		var th uint32
		for j := 0; j < len(v); j++ {
			th = th | (1 << (v[j] - 'a'))
		}
		theMap[th]++
	}

	answer := 0
	for _, v := range theMap {
		answer += v * (v - 1) / 2
	}
	return answer
}

func main() {
	words := []string{"apple", "app", "ppl", "scscs"}
	fmt.Println(similarPairs(words))
}
