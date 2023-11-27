package main

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	qianZhuiHe := make([]int, len(nums)+1)
	for i, v := range nums {
		qianZhuiHe[i+1] = qianZhuiHe[i] + v
	}

	firstNums := make([]int, len(nums)-firstLen+1)
	secondNums := make([]int, len(nums)-secondLen+1)

	for i := firstLen; i < len(qianZhuiHe); i++ {
		firstNums[i-firstLen] = qianZhuiHe[i] - qianZhuiHe[i-firstLen]
	}

	for i := secondLen; i < len(qianZhuiHe); i++ {
		secondNums[i-secondLen] = qianZhuiHe[i] - qianZhuiHe[i-secondLen]
	}

	answer := 0
	for i, v := range firstNums {
		for j, v2 := range secondNums {
			if j+secondLen <= i || i+firstLen <= j {
				answer = max(answer, v+v2)
			}
		}
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}
	firstLen := 4
	secondLen := 3
	fmt.Println(maxSumTwoNoOverlap(nums, firstLen, secondLen))
}
