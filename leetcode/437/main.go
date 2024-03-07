package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {

	answer := 0

	var dfs func(r *TreeNode, sum int)
	dfs = func(r *TreeNode, sum int) {

		if r == nil {
			return
		}

		if sum+r.Val == targetSum {
			answer++
			return
		}

		dfs(r.Left, 0)
		dfs(r.Left, sum+r.Val)

		dfs(r.Right, 0)
		dfs(r.Right, sum+r.Val)
	}

	dfs(root, 0)

	return answer
}

func main() {
	root := &TreeNode{Val: 5}
	node1, node2 := &TreeNode{Val: 4}, &TreeNode{Val: 8}
	root.Left = node1
	root.Right = node2

	node3, node4, node5 := &TreeNode{Val: 11}, &TreeNode{Val: 13}, &TreeNode{Val: 4}
	node1.Left = node3
	node2.Left = node4
	node2.Right = node5

	node3.Left = &TreeNode{Val: 7}
	node3.Right = &TreeNode{Val: 2}

	node5.Left = &TreeNode{Val: 5}
	node5.Right = &TreeNode{Val: 1}

	targetSum := 22
	fmt.Println(pathSum(root, targetSum))
}
