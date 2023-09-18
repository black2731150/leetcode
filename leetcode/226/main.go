package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = tmp

	return root
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 4}

	root.Left = node1
	root.Right = node2

	node2.Right = node3

	fmt.Println(invertTree(root))
}
