package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	mod := int(1e9 + 7)
	n := len(arr)
	answer := 0

	// 创建两个数组来存储每个元素作为最小值时可以延伸的左右范围
	left, right := make([]int, n), make([]int, n)
	stack := []int{} // 单调栈

	// 计算每个元素向左延伸的最大范围
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 清空栈，用于计算每个元素向右延伸的最大范围
	stack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 计算每个元素作为最小值时的贡献
	for i := 0; i < n; i++ {
		answer = (answer + arr[i]*(i-left[i])*(right[i]-i)) % mod
	}

	return answer
}

func main() {
	arr := []int{3, 1, 2, 4}
	fmt.Println(sumSubarrayMins(arr))
}
