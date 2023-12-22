package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}

	x := 0
	for len(queue) > 0 {
		copyQueue := make([]*TreeNode, len(queue))
		copy(copyQueue, queue)
		queue = queue[:0]
		nums := []int{}
		for _, tn := range copyQueue {
			if tn.Left != nil {
				queue = append(queue, tn.Left)
				nums = append(nums, tn.Left.Val)
			}

			if tn.Right != nil {
				queue = append(queue, tn.Right)
				nums = append(nums, tn.Right.Val)
			}
		}

		if x%2 == 0 {
			for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
				nums[i], nums[j] = nums[j], nums[i]
			}

			for i, tn := range queue {
				tn.Val = nums[i]
			}
		}
		x++
	}

	return root
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 8}}}
	fmt.Println(reverseOddLevels(root))
}
