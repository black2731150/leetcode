package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {

	type answer struct {
		min float64
		val int
	}

	ans := answer{
		min: abs(target - float64(root.Val)),
		val: root.Val,
	}

	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}

		if x := abs(target - float64(root.Val)); x <= ans.min {
			if x == ans.min {
				if root.Val < ans.val {
					ans.val = root.Val
				}
			} else {
				ans.min = x
				ans.val = root.Val
			}
		}

		help(root.Left)
		help(root.Right)
	}

	help(root)
	return ans.val
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	target := 1.7868
	fmt.Println(closestValue(root, target))
}
