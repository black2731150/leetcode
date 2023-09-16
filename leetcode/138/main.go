package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	n := 0

	copyListHead := &Node{}
	copyList := copyListHead

	numSlice := []int{}

	for head != nil {

		copyList.Val = head.Val

		n++

		if head.Next != nil {
			copyList.Next = &Node{}
			copyList = copyList.Next
		}

		copyHead := head
		copyHead = copyHead.Random
		j := 0
		for copyHead != nil {
			copyHead = copyHead.Next
			j++
		}
		numSlice = append(numSlice, j)

		head = head.Next
	}

	goodList := copyListHead
	for i := range numSlice {
		copyList := copyListHead
		for j := 0; j < n-numSlice[i]; j++ {
			copyList = copyList.Next
		}
		goodList.Random = copyList
		goodList = goodList.Next
	}

	return copyListHead
}

func main() {
	Node1 := &Node{Val: 7}
	Node2 := &Node{Val: 13}
	Node3 := &Node{Val: 11}
	Node4 := &Node{Val: 10}
	Node5 := &Node{Val: 1}

	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = Node4
	Node4.Next = Node5

	Node1.Random = Node5.Next
	Node2.Random = Node1
	Node3.Random = Node5
	Node4.Random = Node3
	Node5.Random = Node1

	copyList := copyRandomList(Node1)

	for copyList != nil {
		fmt.Print(copyList.Val, "->")
		copyList = copyList.Next
	}
	fmt.Println("nil")
}
