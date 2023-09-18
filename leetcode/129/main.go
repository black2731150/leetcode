package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return xxx(root, 0)

}

func xxx(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}

	sum := 10*num + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return xxx(root.Left, sum) + xxx(root.Right, sum)
}

func main() {
	head := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}

	head.Left = node1
	head.Right = node2

	node2.Left = node3
	node2.Right = node4

	fmt.Println(sumNumbers(head))
}
