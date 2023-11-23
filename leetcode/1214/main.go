package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	var midOrder func(root *TreeNode, nums *[]int)
	midR1, midR2 := []int{}, []int{}

	midOrder = func(root *TreeNode, nums *[]int) {
		if root == nil {
			return
		}
		midOrder(root.Left, nums)
		*nums = append(*nums, root.Val)
		midOrder(root.Right, nums)
	}

	midOrder(root1, &midR1)
	midOrder(root2, &midR2)

	for _, v := range midR1 {
		for _, v2 := range midR2 {
			if v+v2 == target {
				return true
			}
		}
	}
	return false
}

func main() {
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	root2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	target := 7
	fmt.Println(twoSumBSTs(root1, root2, target))
}
