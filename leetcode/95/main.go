package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	answer := []*TreeNode{}

	if start > end {
		answer = append(answer, nil)
		return answer
	}

	for i := start; i <= end; i++ {
		leftTree := generate(start, i-1)
		rightTree := generate(i+1, end)

		for _, l := range leftTree {
			for _, r := range rightTree {
				currentTree := &TreeNode{Val: i}
				currentTree.Left = l
				currentTree.Right = r
				answer = append(answer, currentTree)
			}
		}
	}

	return answer
}

func main() {
	n := 5
	fmt.Println(generateTrees(n))
}
