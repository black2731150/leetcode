package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums []int = make([]int, 0)

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	Root := root

	pOrder(root)

	for i := range nums {
		Root.Val = nums[i]
		Root.Left = nil
		if i != len(nums)-1 {
			Root.Right = &TreeNode{}
			Root = Root.Right
		}
	}

	nums = []int{}

}

func pOrder(root *TreeNode) {
	if root == nil {
		return
	}

	nums = append(nums, root.Val)
	pOrder(root.Left)
	pOrder(root.Right)
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

	flatten(head)

	fmt.Println(nums)
}
