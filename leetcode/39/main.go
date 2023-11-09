package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start, target int)
	backtrack = func(start, target int) {
		if target == 0 {
			// 找到了一个满足条件的组合，将其添加到结果中
			result = append(result, append([]int{}, current...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				// 剪枝：当前候选数大于目标值，不可能满足条件，跳过
				continue
			}

			// 剪枝：去除重复元素，避免重复组合
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			current = append(current, candidates[i])
			backtrack(i, target-candidates[i])
			current = current[:len(current)-1]
		}
	}

	// 开始回溯搜索
	backtrack(0, target)

	return result
}

func main() {
	canditales := []int{1, 2, 3, 4}
	target := 7
	fmt.Println(combinationSum(canditales, target))
}
