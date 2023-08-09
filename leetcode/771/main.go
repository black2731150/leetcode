package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[byte]bool, len(jewels))
	for i := 0; i < len(jewels); i++ {
		jewelsMap[jewels[i]] = true
	}

	answer := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := jewelsMap[stones[i]]; ok {
			answer++
		}
	}
	return answer
}

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
}
