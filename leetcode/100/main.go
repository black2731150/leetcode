package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if q == nil && p == nil {
		return true
	}

	if q == nil || p == nil {
		return false
	}

	return q.Val == p.Val && isSameTree(p.Left, q.Left) && isSameTree(q.Right, p.Right)
}

func main() {
	q := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}

	q.Left = node1
	q.Right = node2

	fmt.Println(isSameTree(q, q))
}
