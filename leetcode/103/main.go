package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	order := [][]int{}

	qurre := []*TreeNode{root}

	for len(qurre) > 0 {

		copyQurre := make([]*TreeNode, len(qurre))
		copy(copyQurre, qurre)

		qurre = []*TreeNode{}

		one := []int{}
		for _, xaxaxa := range copyQurre {
			val := xaxaxa.Val
			one = append(one, val)

			if xaxaxa.Left != nil {
				qurre = append(qurre, xaxaxa.Left)
			}

			if xaxaxa.Right != nil {
				qurre = append(qurre, xaxaxa.Right)
			}
		}
		order = append(order, one)
	}

	for i := range order {

		if i%2 == 1 {
			for i2, j := 0, len(order[i])-1; i2 < j; i2, j = i2+1, j-1 {
				order[i][i2], order[i][j] = order[i][j], order[i][i2]
			}
		}

	}

	return order
}

func main() {
	head := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}

	head.Left = node1
	head.Right = node2

	node2.Left = node3
	node2.Right = node4

	fmt.Println(zigzagLevelOrder(head))
}
