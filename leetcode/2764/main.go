package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isPreorder(nodes [][]int) bool {

	theMap := make(map[int]*TreeNode)

	for _, v := range nodes {
		id := v[0]
		theMap[id] = &TreeNode{Val: id}
	}

	var root *TreeNode
	for _, v := range nodes {
		id := v[0]
		parentId := v[1]

		if parentId == -1 {
			root = theMap[id]
		} else {
			parent := theMap[parentId]
			if parent.Left == nil {
				parent.Left = theMap[id]
				continue
			}

			if parent.Right == nil {
				parent.Right = theMap[id]
				continue
			}
		}
	}

	preOrderNums := []int{}
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		preOrderNums = append(preOrderNums, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)

	fmt.Println(preOrderNums)

	for i, v := range nodes {
		if preOrderNums[i] != v[0] {
			return false
		}
	}
	return true
}

func main() {
	nodes := [][]int{{0, -1}, {1, 0}, {2, 0}, {3, 1}, {4, 2}}
	fmt.Println(isPreorder(nodes))
}
