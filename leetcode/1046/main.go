package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)

	for len(stones) > 1 {
		max := stones[len(stones)-1]
		min := stones[len(stones)-2]

		if max == min {
			stones = stones[:len(stones)-2]
		} else {
			stones = stones[:len(stones)-2]
			stones = append(stones, max-min)
		}

		sort.Ints(stones)
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}

func main() {
	stones := []int{2, 2}
	fmt.Println(lastStoneWeight(stones))
}
