package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	var result []int
	Bianli(root, &result)
	return result
}

func Bianli(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	Bianli(root.Left, result)
	*result = append(*result, root.Val)
	Bianli(root.Right, result)
}
