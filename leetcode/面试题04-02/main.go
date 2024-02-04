package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	v := nums[len(nums)/2]

	answer := &TreeNode{Val: v}
	answer.Left = sortedArrayToBST(nums[:len(nums)/2])
	answer.Right = sortedArrayToBST(nums[len(nums)/2+1:])

	return answer
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums))
}
