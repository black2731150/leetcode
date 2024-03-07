package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	L := 0
	for i, v := range postorder {
		if v == preorder[1] {
			L = i + 1
			break
		}
	}

	root.Left = constructFromPrePost(preorder[1:L+1], postorder[:L])
	root.Right = constructFromPrePost(preorder[L+1:], postorder[L:len(postorder)-1])

	return root
}

func main() {
	// Example usage
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	constructFromPrePost(preorder, postorder)

}
