package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestBSTSubtree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	var midOrder func(root *TreeNode)
	midorder := []int{}
	midOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		midOrder(root.Left)
		midorder = append(midorder, root.Val)
		midOrder(root.Right)
	}

	midOrder(root)

	if isBTS(midorder) {
		return len(midorder)
	} else {
		return max(largestBSTSubtree(root.Left), largestBSTSubtree(root.Right))
	}
}

func isBTS(nums []int) bool {
	n := len(nums)
	if n == 1 || n == 0 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if nums[i+1] >= nums[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
