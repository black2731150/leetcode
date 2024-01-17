package main

import "fmt"

type NumArray struct {
	adds []int
}

func Constructor(nums []int) NumArray {

	adds := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		adds[i] = adds[i-1] + nums[i-1]
	}

	return NumArray{
		adds: adds,
	}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.adds[right+1] - n.adds[left]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(2, 8))
}
