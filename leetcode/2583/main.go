package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return 0
	}

	oneLevel := []*TreeNode{root}
	answers := []int{}

	for len(oneLevel) != 0 {
		one := make([]*TreeNode, len(oneLevel))
		copy(one, oneLevel)
		oneLevel = oneLevel[:0]
		x := 0
		for _, tn := range one {
			x = x + tn.Val
			if tn.Left != nil {
				oneLevel = append(oneLevel, tn.Left)
			}

			if tn.Right != nil {
				oneLevel = append(oneLevel, tn.Right)
			}
		}
		answers = append(answers, x)
	}

	sort.Ints(answers)
	if len(answers) < k {
		return -1
	}
	return int64(answers[len(answers)-k])
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	k := 2
	fmt.Println(kthLargestLevelSum(root, k))
}
