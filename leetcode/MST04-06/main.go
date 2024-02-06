package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	nums := []*TreeNode{}

	var help func(r *TreeNode)
	help = func(r *TreeNode) {
		if r == nil {
			return
		}

		help(r.Left)
		nums = append(nums, r)
		help(r.Right)
	}

	help(root)

	for i, v := range nums {
		if v == p {
			if i+1 == len(nums) {
				return nil
			} else {
				return nums[i+1]
			}
		}
	}
	return nil
}

func main() {
	p := &TreeNode{Val: 1}
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: p}}
	inorderSuccessor(tree, p)
}
