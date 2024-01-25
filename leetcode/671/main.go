package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {

	min := root.Val
	smin := 1 << 31

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		if r.Val > min && r.Val <= smin {
			smin = r.Val
			return
		}

		dfs(r.Left)
		dfs(r.Right)
	}

	dfs(root)

	if smin == 1<<31 {
		return -1
	}
	return smin
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	fmt.Println(findSecondMinimumValue(root))
}
