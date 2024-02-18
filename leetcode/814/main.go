package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var dfs func(r *TreeNode) bool
	dfs = func(r *TreeNode) bool {
		if r == nil {
			return false
		}

		if r.Val == 1 {
			return true
		}

		return dfs(r.Left) || dfs(r.Right)
	}

	if !dfs(root) {
		return nil
	}

	var order func(r *TreeNode)
	order = func(r *TreeNode) {
		if r == nil {
			return
		}

		order(r.Left)

		if !dfs(r.Left) {
			r.Left = nil
		}

		if !dfs(r.Right) {
			r.Right = nil
		}

		order(r.Right)
	}
	order(root)

	return root
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 0}}
	pruneTree(root)
}
