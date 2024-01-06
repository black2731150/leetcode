package main

import "fmt"

func combine(n int, k int) [][]int {
	answer := [][]int{}
	current := []int{}

	var help func(start int)
	help = func(start int) {
		if len(current) == k {
			combination := make([]int, k)
			copy(combination, current)
			answer = append(answer, combination)
			return
		}

		for i := start; i <= n; i++ {
			current = append(current, i)       // 添加当前数字
			help(i + 1)                        // 递归下一层
			current = current[:len(current)-1] // 回溯，移除最后一个数字
		}
	}

	help(1)
	return answer
}

func main() {
	fmt.Println(combine(4, 2))
}
