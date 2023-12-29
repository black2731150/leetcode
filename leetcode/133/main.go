package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	theMap := make(map[int][]*Node)

	queue := []*Node{node}
	for len(queue) > 0 {
		one := queue[0]
		queue = queue[1:]
		theMap[one.Val] = append([]*Node{})
	}
}
