package main

import "fmt"

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)

	if n == 0 {
		return nil
	}

	if n == 1 && len(grid[0]) == 1 {
		return &Node{
			Val:    change(grid[0][0]),
			IsLeaf: true,
		}
	}

	halfN := n / 2
	// 创建四个子切片
	topLeft := make([][]int, halfN)
	topRight := make([][]int, halfN)
	bottomLeft := make([][]int, halfN)
	bottomRight := make([][]int, halfN)

	for i := 0; i < halfN; i++ {
		topLeft[i] = grid[i][:halfN]
		topRight[i] = grid[i][halfN:]
		bottomLeft[i] = grid[i+halfN][:halfN]
		bottomRight[i] = grid[i+halfN][halfN:]
	}

	x := grid[0][0]
	for _, v := range grid {
		for _, v2 := range v {
			if v2 == x {
				continue
			} else {
				return &Node{
					IsLeaf:      false,
					Val:         false,
					TopLeft:     construct(topLeft),
					TopRight:    construct(topRight),
					BottomLeft:  construct(bottomLeft),
					BottomRight: construct(bottomRight),
				}
			}
		}

	}
	return &Node{
		IsLeaf: true,
		Val:    change(x),
	}
}

func change(n int) bool {
	return n == 1
}

func main() {
	gird := [][]int{
		{0, 1}, {1, 0},
	}

	fmt.Println(construct(gird))
}
