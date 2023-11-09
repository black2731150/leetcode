package main

import "fmt"

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
	fmt.Println(midorder)

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

// func largestBSTSubtree(root *TreeNode) int {
//     ret, _, _, _ := ls(root)
//     return ret
// }

// func ls(root *TreeNode) (int, int, int, bool) {
//     if root == nil {
//         return 0, -1 << 31, 1 << 31, true
//     }
//     if root.Left == nil && root.Right == nil {
//         return 1, root.Val, root.Val, true
//     }

//     left, lmin, lmax, ok1 := ls(root.Left)
//     right, rmin, rmax, ok2 := ls(root.Right)
//     if !ok1 || !ok2 {
//         return max(left, right), 0, 0, false
//     }
//     if root.Left != nil && lmax >= root.Val {
//         return max(left, right), 0, 0, false
//     }
//     if root.Right != nil && rmin <= root.Val {
//         return max(left, right), 0, 0, false
//     }
//     if root.Left == nil {
//         return 1 + right, root.Val, rmax, true
//     }
//     if root.Right == nil {
//         return 1 + left, lmin, root.Val, true
//     }
//     return 1 + left + right, lmin, rmax, true
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

func main() {
	root := &TreeNode{Val: 4}

	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 7}

	node4 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 5}

	node7 := &TreeNode{Val: 2}
	node8 := &TreeNode{Val: 1}

	root.Left = node2
	root.Right = node3

	node2.Left = node4
	node2.Right = node5
	node3.Left = node6

	node4.Left = node7
	node7.Left = node8

	fmt.Println(largestBSTSubtree(root))

}
