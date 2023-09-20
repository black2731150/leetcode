package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	levelOrder := [][]*Node{}

	if root == nil {
		return nil
	}

	qurre := []*Node{root}

	for len(qurre) != 0 {
		copyQurre := make([]*Node, len(qurre))
		copy(copyQurre, qurre)
		qurre = []*Node{}

		levelOrder = append(levelOrder, copyQurre)
		for i := range copyQurre {
			if copyQurre[i].Left != nil {
				qurre = append(qurre, copyQurre[i].Left)
			}

			if copyQurre[i].Right != nil {
				qurre = append(qurre, copyQurre[i].Right)
			}

		}
	}

	for i := 0; i < len(levelOrder); i++ {
		for j := 0; j < len(levelOrder[i])-1; j++ {
			levelOrder[i][j].Next = levelOrder[i][j+1]
		}
	}

	return root
}

func main() {
	head := &Node{Val: 3}
	node1 := &Node{Val: 9}
	node2 := &Node{Val: 20}
	node3 := &Node{Val: 15}
	node4 := &Node{Val: 7}

	head.Left = node1
	head.Right = node2

	node2.Left = node3
	node2.Right = node4

	fmt.Println(connect(head))
}
