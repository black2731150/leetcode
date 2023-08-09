package main

type MyHashSet struct {
	HashSet map[int]bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		HashSet: make(map[int]bool),
	}
}

func (h *MyHashSet) Add(key int) {
	h.HashSet[key] = true
}

func (h *MyHashSet) Remove(key int) {
	delete(h.HashSet, key)
}

func (h *MyHashSet) Contains(key int) bool {
	_, find := h.HashSet[key]
	return find
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
