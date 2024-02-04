package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right == nil {
		return left
	}

	if right != nil && left == nil {
		return right
	}

	if left != nil && right != nil {
		return root
	}

	return nil
}

func main() {
	p := &TreeNode{Val: 1}
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: p}}
	lowestCommonAncestor(tree, p, p)
}
