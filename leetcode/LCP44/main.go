package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numColor(root *TreeNode) int {

	var theMap map[int]bool = map[int]bool{}
	print(root, theMap)
	answer := 0
	for _, V := range theMap {
		if V {
			answer++
		}
	}
	return answer
}

func print(root *TreeNode, themap map[int]bool) {
	if root == nil {
		return
	}
	print(root.Left, themap)
	themap[root.Val] = true
	print(root.Right, themap)
}
