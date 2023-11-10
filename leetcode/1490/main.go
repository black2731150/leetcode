package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	answer := &Node{}
	answer.Val = root.Val

	answer.Children = make([]*Node, len(root.Children))
	for i, n := range root.Children {
		answer.Children[i] = cloneTree(n)
	}

	return answer
}

func main() {
	root := &Node{Val: 1}

	node1 := &Node{Val: 2}
	node2 := &Node{Val: 3}
	root.Children = append(root.Children, node1)
	root.Children = append(root.Children, node2)

	fmt.Println(cloneTree(root))
}
