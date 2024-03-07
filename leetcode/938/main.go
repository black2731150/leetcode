package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		dfs(r.Left)

		if r.Val <= high && r.Val >= low {
			sum += r.Val
		}

		dfs(r.Right)
	}

	dfs(root)
	return sum
}
func main() {
	rangeSumBST(nil, 2, 3)
}
