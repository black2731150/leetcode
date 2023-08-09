package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}

func main() {
	fmt.Println(checkTree(nil))
}
