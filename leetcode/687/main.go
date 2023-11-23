package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	answer := 0

	var work func(root *TreeNode) int
	work = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := work(root.Left)
		right := work(root.Right)

		leftNum, rightNum := 0, 0
		if root.Left != nil && root.Val == root.Left.Val {
			leftNum = left + 1
		}

		if root.Right != nil && root.Val == root.Right.Val {
			rightNum = right + 1
		}

		answer = max(answer, leftNum+rightNum)
		return max(leftNum, rightNum)
	}

	work(root)
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}}
	fmt.Println(longestUnivaluePath(root))
}
