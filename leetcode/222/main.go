package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
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

	answer := 0

	for i := range order {
		answer = answer + len(order[i])
	}
	return answer
}
