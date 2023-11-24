package main

import (
	"fmt"
	"sort"
)

type hhh struct {
	num    int
	weight int
}

func getKth(lo int, hi int, k int) int {
	theHHHs := []hhh{}
	for i := lo; i <= hi; i++ {
		weight := getWeight(i)
		theHHHs = append(theHHHs, hhh{
			num:    i,
			weight: weight,
		})
	}

	sort.Slice(theHHHs, func(i, j int) bool {
		if theHHHs[i].weight == theHHHs[j].weight {
			return theHHHs[i].num < theHHHs[j].num
		} else {
			return theHHHs[i].weight < theHHHs[j].weight
		}
	})

	return theHHHs[k-1].num
}

func getWeight(num int) int {
	if num == 1 {
		return 0
	}

	answer := 0

	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}
		answer++
	}
	return answer
}

func main() {
	lo := 12
	hi := 15
	k := 2
	fmt.Println(getKth(lo, hi, k))
}
