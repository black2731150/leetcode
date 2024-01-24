package main

type NestedInteger struct {
}

// Return true if n NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool

// Return the single integer that n NestedInteger holds, if it holds a single integer
// The result is undefined if n NestedInteger holds a nested list
// So before calling n method, you should have a check
func (n NestedInteger) GetInteger() int

// Set n NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int)

// Set n NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger)

// Return the nested list that n NestedInteger holds, if it holds a nested list
// The list length is zero if n NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger

type NestedIterator struct {
	index int
	nums  []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		index: 0,
		nums:  help(nestedList),
	}
}

func help(nestedList []*NestedInteger) []int {
	nums := []int{}
	for _, ni := range nestedList {
		if ni.IsInteger() {
			nums = append(nums, ni.GetInteger())
		} else {
			nums = append(nums, help(ni.GetList())...)
		}
	}
	return nums
}

func (n *NestedIterator) Next() int {
	x := n.nums[n.index]
	n.index++
	return x
}

func (n *NestedIterator) HasNext() bool {
	return n.index != len(n.nums)
}
