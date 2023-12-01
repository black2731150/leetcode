package main

import "fmt"

func prevPermOpt1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// 从后向前找到第一个减少的位置
	i := len(arr) - 2
	for i >= 0 && arr[i] <= arr[i+1] {
		i--
	}

	if i >= 0 {
		// 从后向前找到第一个小于arr[i]的数
		j := len(arr) - 1
		for arr[j] >= arr[i] {
			j--
		}
		// 跳过重复的数字
		for j > 0 && arr[j] == arr[j-1] {
			j--
		}
		// 交换
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(prevPermOpt1(arr))
}
