package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	answer := [][]int{}

	var dfs func(r *TreeNode, path []int, sum int)
	dfs = func(r *TreeNode, path []int, sum int) {

		if r == nil {
			return
		}

		sum = sum - r.Val
		path = append(path, r.Val)

		if sum == 0 && r.Left == nil && r.Right == nil {
			cp := make([]int, len(path))
			copy(cp, path)
			answer = append(answer, cp)
			return
		}

		n := len(path)
		dfs(r.Left, path, sum)
		path = path[:n]
		dfs(r.Right, path, sum)
	}

	dfs(root, []int{}, targetSum)

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
