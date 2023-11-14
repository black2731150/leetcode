package main

import (
	"fmt"
	"sort"
)

type Map struct {
	k   int
	val int
}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)

	allMap := make(map[int]int)
	for _, v := range barcodes {
		allMap[v]++
	}

	AllMap := []Map{}
	for i, v := range allMap {
		AllMap = append(AllMap, Map{k: i, val: v})
	}

	sort.Slice(AllMap, func(i, j int) bool {
		return AllMap[i].val > AllMap[j].val
	})

	answer := make([]int, n)
	answerIndex := 0

	for i := 0; i < len(AllMap); i++ {
		x := AllMap[i]
		for x.val != 0 {
			answer[answerIndex] = x.k
			x.val--
			answerIndex += 2
			if answerIndex >= n {
				answerIndex = 1
			}
		}
	}
	return answer
}

func main() {
	barcodes := []int{1, 1, 1, 2, 2, 2}
	fmt.Println(rearrangeBarcodes(barcodes))
}
