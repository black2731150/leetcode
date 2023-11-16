package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] <= pairs[j][1]
	})

	answer := 0
	end := -1000000001
	for i := 0; i < len(pairs); i++ {
		if pairs[i][0] > end {
			answer++
			end = pairs[i][1]
		}
	}
	return answer
}

func main() {
	paris := [][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}}
	fmt.Println(findLongestChain(paris))
}
