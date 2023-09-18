package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	val := preorder[0]
	preorder = preorder[1:]
	node := &TreeNode{Val: val}

	inorderValIndex := -1
	for i := range inorder {
		if inorder[i] == val {
			inorderValIndex = i
		}
	}

	if inorderValIndex == -1 {
		return nil
	}

	leftPreorder := preorder[:inorderValIndex]

	var rightPreorder []int
	if len(preorder) == 0 {
		rightPreorder = []int{}
	} else {
		rightPreorder = preorder[inorderValIndex:]
	}

	leftInorder := inorder[:inorderValIndex]
	var rightInorder []int
	if len(inorder) == 0 {
		rightInorder = []int{}
	} else {
		rightInorder = inorder[inorderValIndex+1:]
	}

	node.Left = buildTree(leftPreorder, leftInorder)
	node.Right = buildTree(rightPreorder, rightInorder)

	return node
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

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	answer := buildTree(preorder, inorder)

	fmt.Println(answer)
}
