package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func levelOrderBottom(root *TreeNode) [][]int {

	answer := [][]int{}

	if root == nil {
		return answer
	}

	qurre := []*TreeNode{root}

	for len(qurre) > 0 {
		copyQurre := make([]*TreeNode, len(qurre))
		copy(copyQurre, qurre)

		qurre = []*TreeNode{}

		vals := []int{}

		for _, tn := range copyQurre {
			vals = append(vals, tn.Val)

			if tn.Left != nil {
				qurre = append(qurre, tn.Left)
			}

			if tn.Right != nil {
				qurre = append(qurre, tn.Right)
			}
		}

		answer = append(answer, vals)
	}

	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}

	return answer
}

func main() {
	root := NewTreeNode(1)
	node1 := NewTreeNode(2)
	node2 := NewTreeNode(3)

	node3 := NewTreeNode(4)
	node4 := NewTreeNode(5)

	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4

	fmt.Println(levelOrderBottom(root))
}
