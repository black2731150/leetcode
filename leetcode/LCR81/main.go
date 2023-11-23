package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var answer [][]int

	var work func(nums []int, combination []int, target int)
	work = func(nums []int, combination []int, target int) {
		if target == 0 {
			answer = append(answer, append([]int{}, combination...))
			return
		}
		if target < 0 {
			return
		}

		for i, v := range nums {
			work(nums[i:], append(combination, v), target-v)
		}
	}

	work(candidates, []int{}, target)
	return answer
}

func main() {
	candidates := []int{1, 2, 3, 4, 5, 6}
	taeget := 6
	fmt.Println(combinationSum(candidates, taeget))
}
