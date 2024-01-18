package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(r *TreeNode) (int, int)
	dfs = func(r *TreeNode) (int, int) {
		if r == nil {
			return 0, 0
		}

		leftWithRoot, leftWithOutRoot := dfs(r.Left)
		rightWitRoot, rightWithOutRoot := dfs(r.Right)

		withRoot := r.Val + leftWithOutRoot + rightWithOutRoot
		withoutRoot := max(leftWithRoot, leftWithOutRoot) + max(rightWitRoot, rightWithOutRoot)

		return withRoot, withoutRoot
	}

	withRoot, withoutRoot := dfs(root)
	return max(withRoot, withoutRoot)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}
	fmt.Println(rob(root))
}
