package main

import "fmt"

func giveGem(gem []int, operations [][]int) int {
	max := 0
	min := 1 << 31
	for _, v := range operations {
		giver := v[0]
		geter := v[1]

		gem[geter] = gem[geter] + gem[giver]/2
		gem[giver] = gem[giver] - gem[giver]/2
	}

	for _, v := range gem {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max - min
}

func main() {
	gem := []int{1, 2, 3, 4}
	operations := [][]int{{1, 2}, {2, 3}}
	fmt.Println(giveGem(gem, operations))
}
