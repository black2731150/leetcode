package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func heightOfTree(root *TreeNode) int {

	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		lD, rD := 0, 0
		if r.Left != nil && r.Left.Right != r {
			lD = 1 + dfs(r.Left)
		}

		if r.Right != nil && r.Right.Left != r {
			rD = 1 + dfs(r.Right)
		}

		return max(lD, rD)
	}

	return dfs(root)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(heightOfTree(root))
}
