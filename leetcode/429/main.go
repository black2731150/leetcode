package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	answer := [][]int{}
	if root == nil {
		return answer
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		n := len(queue)
		one := []int{}
		for i := 0; i < n; i++ {
			if queue[i] != nil {
				one = append(one, queue[i].Val)
				queue = append(queue, queue[i].Children...)
			}
		}

		queue = queue[n:]
		answer = append(answer, one)
	}

	return answer
}

func main() {
	root := &Node{Val: 1, Children: []*Node{{Val: 2}, {Val: 3}}}
	fmt.Println(levelOrder(root))
}
