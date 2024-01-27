package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	theMap := make(map[int][]int)

	queue := []*Node{node}
	for len(queue) > 0 {
		one := queue[0]
		queue = queue[1:]
		if _, find := theMap[one.Val]; find {
			continue
		} else {
			theMap[one.Val] = make([]int, 0)
			for _, n := range one.Neighbors {
				theMap[one.Val] = append(theMap[one.Val], n.Val)
				queue = append(queue, n)
			}
		}
	}

	nodes := make([]*Node, len(theMap))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}

	for index, naberses := range theMap {
		for _, v := range naberses {
			nodes[index-1].Neighbors = append(nodes[index-1].Neighbors, nodes[v-1])
		}
	}

	for k := range theMap {
		delete(theMap, k)
	}

	return nodes[0]
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = append(node1.Neighbors, node2, node3)
	node2.Neighbors = append(node2.Neighbors, node1, node4)
	node3.Neighbors = append(node3.Neighbors, node1, node4)
	node4.Neighbors = append(node4.Neighbors, node3, node2)

	cloneGraph(node1)
}
