package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums []int = []int{}

func getMinimumDifference(root *TreeNode) int {
	inorder(root)

	min := abs(nums[0], nums[1])
	for i := 0; i < len(nums)-1; i++ {
		if x := abs(nums[i], nums[i+1]); x < min {
			min = x
		}
	}

	nums = nums[:0]
	return min
}

func abs(x, y int) int {
	if answer := x - y; answer >= 0 {
		return answer
	} else {
		return -answer
	}
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	nums = append(nums, root.Val)
	inorder(root.Right)
}

func main() {
	head := &TreeNode{Val: 20}
	node1 := &TreeNode{Val: 15}
	node2 := &TreeNode{Val: 23}
	node3 := &TreeNode{Val: 21}
	node4 := &TreeNode{Val: 40}

	head.Left = node1
	head.Right = node2

	node2.Left = node3
	node2.Right = node4

	fmt.Println(getMinimumDifference(head))
}
