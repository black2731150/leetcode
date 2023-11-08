package main

import (
	"fmt"
	"sort"
)

type Num struct {
	Index int
	Val   int
}

func maxSubsequence(nums []int, k int) []int {
	Nums := []Num{}
	for i, v := range nums {
		Nums = append(Nums, Num{
			Index: i,
			Val:   v,
		})
	}

	sort.Slice(Nums, func(i, j int) bool {
		return Nums[i].Val < Nums[j].Val
	})

	answer := Nums[len(Nums)-k:]

	sort.Slice(answer, func(i, j int) bool {
		return answer[i].Index < answer[j].Index
	})

	ans := []int{}
	for _, n := range answer {
		ans = append(ans, n.Val)
	}
	return ans
}

func main() {
	nums := []int{1, 2, 4, 5, 6, 7}
	k := 2
	fmt.Println(maxSubsequence(nums, k))
}
