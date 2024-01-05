package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	answer := []string{}
	if root == nil {
		return answer
	}

	var help func(root *TreeNode, str string)
	help = func(root *TreeNode, str string) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			answer = append(answer, fmt.Sprintf("%s->%d", str, root.Val)[2:])
		} else {
			if root.Left != nil {
				help(root.Left, fmt.Sprintf("%s->%d", str, root.Val))
			}
			if root.Right != nil {
				help(root.Right, fmt.Sprintf("%s->%d", str, root.Val))
			}
		}
	}

	help(root, "")
	return answer
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}}
	fmt.Println(binaryTreePaths(root))
}
