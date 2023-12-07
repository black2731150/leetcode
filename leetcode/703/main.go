package main

import "sort"

type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{
		nums: nums,
		k:    k,
	}
}

func (k *KthLargest) Add(val int) int {
	k.nums = append(k.nums, val)
	sort.Ints(k.nums)
	return k.nums[len(k.nums)-k.k]
}

func main() {
	obj := Constructor(2, []int{3, 4, 7, 8})
	obj.Add(2)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
