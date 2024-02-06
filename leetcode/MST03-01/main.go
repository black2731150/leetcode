package main

type TripleInOne struct {
	stack      []int
	stackSize  int
	i0, i1, i2 int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack:     make([]int, stackSize*3),
		stackSize: stackSize,
		i0:        0,
		i1:        stackSize,
		i2:        stackSize * 2,
	}
}

func (t *TripleInOne) Push(stackNum int, value int) {
	switch stackNum {
	case 0:
		if t.i0 < t.stackSize*(0+1) {
			t.stack[t.i0] = value
			t.i0++
		}
	case 1:
		if t.i1 < t.stackSize*(1+1) {
			t.stack[t.i1] = value
			t.i1++
		}
	case 2:
		if t.i2 < t.stackSize*(2+1) {
			t.stack[t.i2] = value
			t.i2++
		}
	}
}

func (t *TripleInOne) Pop(stackNum int) int {
	v := -1
	switch stackNum {
	case 0:
		if t.i0 > t.stackSize*(0) {
			t.i0--
			v = t.stack[t.i0]
		}
	case 1:
		if t.i1 > t.stackSize*(1) {
			t.i1--
			v = t.stack[t.i1]
		}
	case 2:
		if t.i2 > t.stackSize*(2) {
			t.i2--
			v = t.stack[t.i2]
		}
	}

	return v
}

func (t *TripleInOne) Peek(stackNum int) int {
	v := -1
	switch stackNum {
	case 0:
		if t.i0 > t.stackSize*(0) {
			v = t.stack[t.i0-1]
		}
	case 1:
		if t.i1 > t.stackSize*(1) {
			v = t.stack[t.i1-1]
		}
	case 2:
		if t.i2 > t.stackSize*(2) {
			v = t.stack[t.i2-1]
		}
	}

	return v
}

func (t *TripleInOne) IsEmpty(stackNum int) bool {
	switch stackNum {
	case 0:
		if t.i0 == t.stackSize*(0) {
			return true
		}
	case 1:
		if t.i1 == t.stackSize*(1) {
			return true
		}
	case 2:
		if t.i2 == t.stackSize*(2) {
			return true
		}
	}
	return false
}

func main() {

	stackSize := 3
	obj := Constructor(stackSize)
	obj.Push(1, 2)
	obj.Pop(1)
	obj.Peek(1)
	obj.IsEmpty(2)

}
