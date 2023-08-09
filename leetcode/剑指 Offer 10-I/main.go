package main

import "math"

func fib(n int) int {
	list := []int{}
	list = append(list, 0)
	list = append(list, 1)
	for i := 2; i < n+1; i++ {
		list = append(list, (list[i-1]+list[i-2])%(int(math.Pow10(9)+7)))
	}
	return list[n]

}
