package main

type ListNode struct {
	Key   int
	Val   int
	Front *ListNode
	Next  *ListNode
}

type LRUCache struct {
	MaxLen   int
	RealLen  int
	ListHead *ListNode
	ListTail *ListNode
	TheMap   map[int]*ListNode
}

func Constructor(capacity int) LRUCache {
	node := LRUCache{
		MaxLen:   capacity,
		RealLen:  0,
		ListHead: &ListNode{},
		ListTail: &ListNode{},
		TheMap:   make(map[int]*ListNode),
	}
	node.ListHead.Next = node.ListTail
	node.ListTail.Front = node.ListHead
	return node
}

func (LUR *LRUCache) MoveToHead(node *ListNode) {
	LUR.RmoveNode(node)
	LUR.AddToHead(node)
}

func (LUR *LRUCache) AddToHead(node *ListNode) {
	node.Front = LUR.ListHead
	node.Next = LUR.ListHead.Next
	LUR.ListHead.Next.Front = node
	LUR.ListHead.Next = node
}

func (LUR *LRUCache) RmoveNode(node *ListNode) {
	node.Front.Next = node.Next
	node.Next.Front = node.Front
}

func (LUR *LRUCache) RmoveTail() *ListNode {
	node := LUR.ListTail.Front
	LUR.RmoveNode(node)
	return node
}

func (LUR *LRUCache) Get(key int) int {
	if value, find := LUR.TheMap[key]; find {
		node := LUR.TheMap[key]
		LUR.MoveToHead(node)
		return value.Val
	}
	return -1
}

func (LUR *LRUCache) Put(key int, value int) {
	if _, find := LUR.TheMap[key]; find {
		node := LUR.TheMap[key]
		node.Val = value
		LUR.MoveToHead(node)
	} else {
		node := &ListNode{
			Key:   key,
			Val:   value,
			Front: nil,
			Next:  nil,
		}
		LUR.TheMap[key] = node
		LUR.AddToHead(node)
		LUR.RealLen++
		if LUR.RealLen > LUR.MaxLen {
			removed := LUR.RmoveTail()
			delete(LUR.TheMap, removed.Key)
			LUR.RealLen--
		}
	}
}

func main() {
	lurCache := Constructor(2)
	lurCache.Put(1, 1)
	lurCache.Put(2, 2)
	lurCache.Get(1)
	lurCache.Put(3, 3)
}
