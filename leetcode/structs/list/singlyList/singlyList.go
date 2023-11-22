package singlylist

type SinglyListNode struct {
	Val  int
	Next *SinglyListNode
}

// 创建一个新的单项链表
func NewSinglyListNode(val int) *SinglyListNode {
	return &SinglyListNode{
		Val:  val,
		Next: nil,
	}
}

// 给链表的头部添加元素
func (s *SinglyListNode) AddToHead(node *SinglyListNode) *SinglyListNode {
	node.Next = s
	return node
}

// 给链表的末尾添加元素
func (s *SinglyListNode) AddToEnd(node *SinglyListNode) *SinglyListNode {
	head := s
	if s != nil {
		for s.Next != nil {
			s = s.Next
		}
		s.Next = node
		return node
	}
	return head
}

// 获取链表的长度
func (s *SinglyListNode) Lenght() int {
	head := s
	answer := 0
	for head != nil {
		answer++
		head = head.Next
	}
	return answer
}

// 链表中查找一个元素是否存在
func (s *SinglyListNode) Search(val int) bool {
	head := s
	for head != nil {
		if head.Val == val {
			return true
		}
		head = head.Next
	}
	return false
}

// 删除链表头部的元素
func (s *SinglyListNode) DeleteHeadNode() *SinglyListNode {
	if s != nil {
		return s.Next
	}
	return nil
}

// 删除链表尾部的元素
func (s *SinglyListNode) DeleteEndNode() *SinglyListNode {
	head := s
	if s == nil || s.Next == nil {
		return nil
	}

	if s.Next.Next != nil {
		s = s.Next
	}
	s.Next = nil

	return head
}

// 遍历链表返回链表中的值组成的切片
func (s *SinglyListNode) Range() []int {
	answer := []int{}
	if s == nil {
		return answer
	}

	for s != nil {
		answer = append(answer, s.Val)
		s = s.Next
	}
	return answer
}

// 清楚这个链表
func (s *SinglyListNode) Clear() {
	s = nil
}

// 判断这个链表是否是空链表
func (s *SinglyListNode) IsEmpty() bool {
	return s == nil
}

// 反转链表
func (s *SinglyListNode) Reverse() *SinglyListNode {
	if s == nil {
		return nil
	}

	if s.Next == nil {
		return s
	}

}
