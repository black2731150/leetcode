package main

type node struct {
	val  int
	next *node
	prev *node
}

type MyLinkedList struct {
	Nums []*node
	Head *node
	End  *node
}

func Constructor() MyLinkedList {
	head := &node{}
	end := &node{}
	head.next = end
	end.prev = head
	return MyLinkedList{
		Nums: make([]*node, 0),
		Head: head,
		End:  end,
	}
}

func (mll *MyLinkedList) Get(index int) int {
	if index < 0 || index >= len(mll.Nums) {
		return -1
	} else {
		return mll.Nums[index].val
	}
}

func (mll *MyLinkedList) AddAtHead(val int) {
	new := &node{val: val}
	new.prev = mll.Head
	new.next = mll.Head.next
	mll.Head.next.prev = new
	mll.Head.next = new

	mll.Nums = append([]*node{new}, mll.Nums...)
}

func (mll *MyLinkedList) AddAtTail(val int) {
	new := &node{val: val}
	new.prev = mll.End.prev
	new.next = mll.End
	mll.End.prev.next = new
	mll.End.prev = new

	mll.Nums = append(mll.Nums, new)
}

func (mll *MyLinkedList) AddAtIndex(index int, val int) {
	if index > len(mll.Nums) || index < 0 {
		return
	}

	if index == len(mll.Nums) {
		mll.AddAtTail(val)
		return
	}

	if index == 0 {
		mll.AddAtHead(val)
		return
	}

	oldNode := mll.Nums[index]
	new := &node{val: val}

	new.next = oldNode
	new.prev = oldNode.prev
	oldNode.prev = new
	new.prev.next = new

	back := append([]*node{new}, mll.Nums[index:]...)
	mll.Nums = append(mll.Nums[:index], back...)
}

func (mll *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= len(mll.Nums) {
		return
	} else {

		oldNode := mll.Nums[index]
		oldNode.next.prev = oldNode.prev
		oldNode.prev.next = oldNode.next
		oldNode.next = nil
		oldNode.prev = nil

		if index == 0 {
			mll.Nums = mll.Nums[1:]
		} else {
			mll.Nums = append(mll.Nums[:index], mll.Nums[index+1:]...)
		}
	}
}

func main() {
	obj := Constructor()
	obj.Get(1)
	obj.AddAtHead(1)
	obj.AddAtTail(2)
	obj.AddAtIndex(1, 0)
	obj.DeleteAtIndex(1)
}
