package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)

	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[n/2]}

	root.Left = sortedArrayToBST(nums[:n/2])

	if n/2 != 0 {
		root.Right = sortedArrayToBST(nums[(n/2)+1:])
	} else {
		root.Right = sortedArrayToBST([]int{})
	}

	return root
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	tree := sortedArrayToBST(nums)
	fmt.Println(tree)
}
