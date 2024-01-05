package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度为3n，够了，但是可以优化
func removeNodes(head *ListNode) *ListNode {
	copyHead := head
	nums := []int{}
	for copyHead != nil {
		nums = append(nums, copyHead.Val)
		copyHead = copyHead.Next
	}

	maxNums := make([]int, len(nums))
	maxNum := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
		maxNums[i] = maxNum
	}

	answer := &ListNode{}
	answerHead := answer

	i := 0
	for head != nil {
		if head.Val >= maxNums[i] {
			answer.Next = &ListNode{Val: head.Val}
			answer = answer.Next
		}
		i++
		head = head.Next
	}

	return answerHead.Next
}

func main() {
	// 创建测试链表: [1, 3, 2, 4, 3]
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}}

	// 调用 removeNodes 函数
	result := removeNodes(head)

	// 打印结果
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}

//时间复杂度过高
// func removeNodes(head *ListNode) *ListNode {
// 	answer := &ListNode{}
// 	answerHead := answer
// 	for head != nil {
// 		firsVal := head.Val
// 		success := true
// 		copyHead := head
// 		for copyHead.Next != nil {
// 			if copyHead.Next.Val > firsVal {
// 				success = false
// 				break
// 			}
// 			copyHead = copyHead.Next
// 		}

// 		if success {
// 			answer.Next = &ListNode{firsVal, nil}
// 			answer = answer.Next
// 		}

// 		head = head.Next
// 	}

// 	return answerHead.Next
// }
