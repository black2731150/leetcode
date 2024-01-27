package main

import "fmt"

func change(amount int, coins []int) int {

	bp := make([]int, amount+1)
	bp[0] = 0
	bp[1] = 1
	bp[2] = 2

	for i := 1; i < len(bp); i++ {
		for _, v := range coins {
			if i-v >= 0 {
				bp[i] = bp[i-v] + bp[v]
			}
		}
	}

	fmt.Println(bp)

	return bp[len(bp)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
