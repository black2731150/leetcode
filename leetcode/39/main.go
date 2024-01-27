package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	answer := [][]int{}

	var help func(nums []int, index int, tar int)
	help = func(nums []int, index int, tar int) {
		if tar < 0 {
			return
		}

		if tar == 0 {
			copyOne := make([]int, len(nums))
			copy(copyOne, nums)
			answer = append(answer, copyOne)
			return
		} else {
			for i := index; i < len(candidates); i++ {
				if tar < candidates[i] {
					break
				}

				if i > index && candidates[i] == candidates[i-1] {
					continue
				}

				nums = append(nums, candidates[i])
				help(nums, i, tar-candidates[i])
				nums = nums[:len(nums)-1]
			}
		}
	}

	help([]int{}, 0, target)
	return answer
}

func main() {
	canditales := []int{1, 2, 3, 4}
	target := 7
	fmt.Println(combinationSum(canditales, target))
}
