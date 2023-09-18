package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	return check(root, root)

}

func check(q, p *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}

	if q == nil || p == nil {
		return false
	}

	return p.Val == q.Val && check(q.Left, p.Right) && check(q.Right, p.Left)
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 3}

	root.Left = node1
	root.Right = node2

	node1.Left = node3
	node1.Right = node4

	node2.Left = node5
	node2.Right = node6

	fmt.Println(isSymmetric(root))
}
