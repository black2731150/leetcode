package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := nums[0]
	maxIndex := 0
	for i, v := range nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}

	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(constructMaximumBinaryTree(nums))
}
