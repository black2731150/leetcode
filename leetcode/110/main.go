package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(Hight(root.Left)-Hight(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Hight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(Hight(root.Left), Hight(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(isBalanced(root))
}
