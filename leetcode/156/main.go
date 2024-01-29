package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var rootLeft, rootRight *TreeNode

	if root.Left != nil {
		rootLeft = root.Left
		rootLeft = upsideDownBinaryTree(rootLeft)
		rootLeft.Left = rootRight
	}

	if root.Right != nil {
		rootRight = root.Right
		rootRight = upsideDownBinaryTree(rootRight)
		rootLeft.Right = root
	}

	return rootLeft
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(root)
}
