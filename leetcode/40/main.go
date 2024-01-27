package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	answer := [][]int{}

	tmp := []int{}
	var help func(start int, t int)
	help = func(start int, t int) {
		if t < 0 {
			return
		}

		if t == 0 {
			copyNums := make([]int, len(tmp))
			copy(copyNums, tmp)

			answer = append(answer, copyNums)
			return

		}

		for i := start; i < n; i++ {

			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			tmp = append(tmp, candidates[i])
			help(i+1, t-candidates[i])
			tmp = tmp[:len(tmp)-1]
		}
	}
	help(0, target)
	return answer
}

func main() {
	ca := []int{2, 5, 2, 1, 2}
	target := 5
	fmt.Println(combinationSum2(ca, target))
}
