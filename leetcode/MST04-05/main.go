package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return help(root, math.MinInt, math.MaxInt)
}

func help(root *TreeNode, low, up int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= up {
		return false
	}

	return help(root.Left, low, root.Val) && help(root.Right, root.Val, up)
}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	isValidBST(tree)
}
