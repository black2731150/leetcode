package main

import "fmt"

func sumIndicesWithKSetBits(nums []int, k int) int {
	answer := 0
	for i, v := range nums {
		if num0fOne(i) == k {
			answer += v
		}
	}
	return answer
}

func num0fOne(x int) int {
	count := 0
	for x > 0 {
		x &= x - 1 // 清除最低位的1
		count++
	}
	return count
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 1
	fmt.Println(sumIndicesWithKSetBits(nums, k))
}
