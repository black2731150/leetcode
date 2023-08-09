package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var answer []int = []int{}
	if root == nil {
		return answer
	}
	print(root, &answer)

	return answer
}

func print(root *TreeNode, answer *[]int) {
	if root == nil {
		return
	}

	*answer = append(*answer, root.Val)
	print(root.Left, answer)
	print(root.Right, answer)
}

func main() {
	fmt.Println(preorderTraversal(nil))
}
