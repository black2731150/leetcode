package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		return &TreeNode{Val: val, Left: root}
	}

	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	root := constructMaximumBinaryTree(nums)
	fmt.Println(insertIntoMaxTree(root, 8))
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
