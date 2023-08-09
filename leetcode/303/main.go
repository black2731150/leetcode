package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	answer := 0
	i, j := 0, 0
	for i, j = left, right; i < j; i, j = i+1, j-1 {
		answer = answer + this.nums[i] + this.nums[j]
	}
	answer = answer + this.nums[i]
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(2, 8))
}
