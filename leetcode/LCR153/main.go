package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathTarget(root *TreeNode, target int) [][]int {
	var answer [][]int = [][]int{}

	var work func(root *TreeNode, target int, one []int)
	work = func(root *TreeNode, target int, one []int) {
		if root == nil {
			return
		}

		val := root.Val
		one1 := append([]int{}, one...)
		one1 = append(one1, val)

		if root.Left == nil && root.Right == nil {
			if target-val == 0 {
				answer = append(answer, one1)
			}
		} else {
			work(root.Left, target-val, one1)
			work(root.Right, target-val, one1)
		}
	}

	work(root, target, []int{})

	return answer
}

func main() {
	root := TreeNode{Val: 5}
	node1 := TreeNode{Val: 4}
	node2 := TreeNode{Val: 8}
	root.Left = &node1
	root.Right = &node2

	node3 := TreeNode{Val: 11}
	node4 := TreeNode{Val: 13}
	node5 := TreeNode{Val: 4}

	node1.Left = &node3
	node2.Left = &node4
	node2.Right = &node5

	node6 := TreeNode{Val: 7}
	node7 := TreeNode{Val: 2}
	node8 := TreeNode{Val: 5}
	node9 := TreeNode{Val: 1}

	node3.Left = &node6
	node3.Right = &node7
	node5.Left = &node8
	node5.Right = &node9

	fmt.Println(pathTarget(&root, 22))
}
