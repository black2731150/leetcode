package main

import (
	"fmt"
	"math"
)

func maximumSubarraySum(nums []int, k int) int64 {
	v2Index := make(map[int][]int)

	qianzui := make([]int64, len(nums)+1)
	answer := int64(math.MinInt64)
	for i, v := range nums {
		qianzui[i+1] = qianzui[i] + int64(v)
		if indexs, find := v2Index[v-k]; find {
			for _, index := range indexs {
				answer = max(answer, qianzui[i+1]-qianzui[index])
			}
		}

		if indexs, find := v2Index[v+k]; find {
			for _, index := range indexs {
				answer = max(answer, qianzui[i+1]-qianzui[index])
			}
		}
		v2Index[v] = append(v2Index[v], i)
	}

	if answer == int64(math.MinInt64) {
		return 0
	}

	return answer
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 5}
	k := 3
	fmt.Println(maximumSubarraySum(nums, k))
}
