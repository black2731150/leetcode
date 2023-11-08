package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sort := []int{}
	noods := []*TreeNode{}
	all := 0
	var getOrder func(root *TreeNode)
	getOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		getOrder(root.Left)
		sort = append(sort, root.Val)
		all = all + root.Val
		noods = append(noods, root)
		getOrder(root.Right)
	}
	getOrder(root)

	n := len(sort)
	left := 0
	right := all
	allSum := []int{}
	for i := 0; i < n; i++ {
		allSum = append(allSum, right-left)
		left = left + sort[i]
		noods[i].Val = allSum[i]
	}

	return root
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

	fmt.Println(bstToGst(root))
}
