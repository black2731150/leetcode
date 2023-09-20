package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	nums := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		nums = append(nums, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}

func main() {
	head := &TreeNode{Val: 20}
	node1 := &TreeNode{Val: 15}
	node2 := &TreeNode{Val: 23}
	node3 := &TreeNode{Val: 21}
	node4 := &TreeNode{Val: 40}

	head.Left = node1
	head.Right = node2

	node2.Left = node3
	node2.Right = node4

	fmt.Println(isValidBST(head))
}
