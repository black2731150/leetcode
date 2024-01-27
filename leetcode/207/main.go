package main

import "fmt"

type Node struct {
	Val int
	In  map[int]struct{}
	Out map[int]struct{}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
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

	for {
		found := false
		for i, node := range nodes {
			if len(node.In) == 0 && len(node.Out) > 0 {
				found = true
				for k := range node.Out {
					delete(nodes[k].In, i)
				}
				nodes[i].Out = make(map[int]struct{})
				break
			}
		}
		if !found {
			break
		}
	}

	answer := 0
	for _, n := range nodes {
		if len(n.In) == 0 {
			answer++
		}
	}

	return answer == numCourses
}

func main() {
	numsC := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numsC, prerequisites))
}
