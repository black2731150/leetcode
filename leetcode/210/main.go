package main

import "fmt"

type Node struct {
	Val        int
	In         map[int]struct{}
	Out        map[int]struct{}
	OutVisited bool
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	nodes := make([]Node, numCourses)
	for i := range nodes {
		nodes[i].Val = i
		nodes[i].Out = make(map[int]struct{})
		nodes[i].In = make(map[int]struct{})
	}

	for _, one := range prerequisites {
		start := one[1]
		end := one[0]

		nodes[start].Out[end] = struct{}{}
		nodes[end].In[start] = struct{}{}
	}

	stack := []int{}
	for {
		found := false
		for i, node := range nodes {
			if len(node.In) == 0 && !node.OutVisited {
				found = true
				for k := range node.Out {
					delete(nodes[k].In, i)
				}
				nodes[i].OutVisited = true
				stack = append(stack, i)
				break
			}
		}
		if !found {
			break
		}
	}

	for i, node := range nodes {
		if len(node.In) == 0 && !node.OutVisited {
			stack = append(stack, i)
		}
	}

	if len(stack) == len(nodes) {
		return stack
	} else {
		return []int{}
	}
}

func main() {
	numsC := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(findOrder(numsC, prerequisites))
}
