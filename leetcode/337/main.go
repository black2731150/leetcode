package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(r *TreeNode, parentVal, grandVal int) int
	dfs = func(r *TreeNode, parentVal, grandVal int) int {
		if r == nil {
			return 0
		}
		left := dfs(r.Left, r.Val, parentVal)
		right := dfs(r.Right, r.Val, parentVal)

		return r.Val + max(left, right)
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
