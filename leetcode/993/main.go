package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		fmt.Println(queue)
		n := len(queue)
		findx, findy := false, false
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Val == x {
				findx = true
			}

			if node.Val == y {
				findy = true
			}
			if node.Left != nil && node.Right != nil {
				if (node.Left.Val == x && node.Right.Val == y) || (node.Left.Val == y && node.Right.Val == x) {
					return false
				} else {
					queue = append(queue, node.Left)
					queue = append(queue, node.Right)
				}
			} else {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}

		queue = queue[n:]
		if findx && findy {
			return true
		}

	}
	return false
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	x := 2
	y := 3
	fmt.Println(isCousins(root, x, y))
}
