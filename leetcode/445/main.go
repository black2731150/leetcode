package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := []int{}
	list2 := []int{}

	for l1 != nil {
		list1 = append(list1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		list2 = append(list2, l2.Val)
		l2 = l2.Next
	}

	n1, n2 := len(list1), len(list2)
	fmt.Println("n1=", n1, "  n2=", n2)
	if n1 > n2 {
		x := []int{}
		for i := 0; i < n1-n2; i++ {
			x = append(x, 0)
		}
		list2 = append(x, list2...)
	} else {
		x := []int{}
		for i := 0; i < n2-n1; i++ {
			x = append(x, 0)
		}
		list1 = append(x, list1...)
	}

	fmt.Println(len(list1), "", len(list2))

	n := len(list1)
	count := 0
	answerList := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		x := list1[i] + list2[i] + count
		count = 0
		if x >= 10 {
			x = x - 10
			count++
		}

		answerList[i] = x
	}

	if count == 1 {
		answerList = append([]int{1}, answerList...)
		n = n + 1
	}

	fmt.Println(answerList)
	answerHead := &ListNode{}
	copyAnswerHead := answerHead
	for i, v := range answerList {
		answerHead.Val = v
		if i != n-1 {
			answerHead.Next = &ListNode{}
			answerHead = answerHead.Next
		}
	}

	return copyAnswerHead
}

func main() {
	list1 := &ListNode{Val: 5}
	list2 := &ListNode{Val: 5}
	fmt.Println(addTwoNumbers(list1, list2))
}
