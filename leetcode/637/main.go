package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
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

	fmt.Println(order)

	answer := []float64{}
	for i := range order {
		var one float64
		for j := range order[i] {
			one = one + float64(order[i][j])
		}
		answer = append(answer, one/float64(len(order[i])))
	}

	return answer
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

	fmt.Println(averageOfLevels(head))
}
