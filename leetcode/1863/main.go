package main

import "fmt"

func subsetXORSum(nums []int) int {
	ziji := ziji(nums)
	answer := 0
	for _, v := range ziji {
		tmp := 0
		for _, x := range v {
			tmp = xor(tmp, x)
		}
		answer = answer + tmp
	}
	return answer
}

func xor(x, y int) int {
	return x ^ y
}

func ziji(nums []int) [][]int {
	var result [][]int

	// 开始回溯生成子集
	backtrack(nums, []int{}, &result, 0)

	return result
}

// 回溯法生成子集
func backtrack(nums, subset []int, result *[][]int, start int) {
	// 将当前子集加入结果
	*result = append(*result, append([]int{}, subset...))

	// 从start开始向后遍历元素
	for i := start; i < len(nums); i++ {
		// 将当前元素加入子集
		subset = append(subset, nums[i])

		// 继续回溯生成子集
		backtrack(nums, subset, result, i+1)

		// 回溯，将最后一个元素从子集中移除，继续生成其他子集
		subset = subset[:len(subset)-1]
	}
}

func main() {
	nums := []int{1, 3}
	fmt.Println(subsetXORSum(nums))

}
