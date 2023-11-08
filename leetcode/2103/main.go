package main

import "fmt"

func countPoints(rings string) int {
	theMap := make(map[byte]uint8, 10)

	towN := len(rings)

	for i := 0; i < towN; i = i + 2 {
		switch rings[i] {
		case 'B':
			theMap[rings[i+1]] = theMap[rings[i+1]] | (1 << 0)
		case 'R':
			theMap[rings[i+1]] = theMap[rings[i+1]] | (1 << 1)
		case 'G':
			theMap[rings[i+1]] = theMap[rings[i+1]] | (1 << 2)
		}
	}

	answer := 0
	for _, v := range theMap {
		if v == 7 {
			answer++
		}
	}
	return answer
}

func main() {
	rings := "B0B6G0R6R0R6G9"
	fmt.Println(countPoints(rings))
}
