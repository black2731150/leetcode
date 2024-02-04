package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	answer := 0
	var dfs func(r *TreeNode, target int)
	dfs = func(r *TreeNode, target int) {
		if r == nil {
			return
		}

		target = target - r.Val
		if target == 0 {
			answer++
		}

		dfs(r.Left, target)
		dfs(r.Right, target)

	}
	var order func(r *TreeNode)
	order = func(r *TreeNode) {
		if r == nil {
			return
		}

		dfs(r, sum)
		order(r.Left)
		order(r.Right)
	}

	order(root)

	return answer
}
