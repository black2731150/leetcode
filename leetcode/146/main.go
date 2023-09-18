package main

type List struct {
	Key  int
	Val  int
	Next *List
}

type LRUCache struct {
	head *List
	max  int
	num  int
	last *List
}

func Constructor(capacity int) LRUCache {

	head := &List{}

	return LRUCache{
		head: head,
		max:  capacity,
		num:  0,
		last: head,
	}
}

func (this *LRUCache) Get(key int) int {
	newHead := &List{Val: 0, Next: this.head}
	Head := newHead
	for ; newHead.Next != nil; newHead = newHead.Next {
		if newHead.Next.Key == key {

			val := newHead.Next.Val

			this.head = Head

			newHead.Next = newHead.Next.Next

			return val
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	headCopy := this.head
	for ; headCopy != nil; headCopy = headCopy.Next {
		if headCopy.Key == key {
			headCopy.Val = value
			return
		}
	}

	if this.num < this.max {

		if this.num == 0 {
			this.last.Key = key
			this.last.Val = value
		} else {
			this.last.Next = &List{
				Key: key,
				Val: value,
			}
			this.last = this.last.Next
		}

		this.num++
	} else {
		this.last.Key = key
		this.last.Val = value
	}
}

func main() {
	lurCache := Constructor(2)
	lurCache.Put(1, 1)
	lurCache.Put(2, 2)
	lurCache.Get(1)
	lurCache.Put(3, 3)
}
