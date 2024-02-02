package main

type SortedStack struct {
	stack []int
	tmp   []int
}

func Constructor() SortedStack {
	return SortedStack{
		stack: []int{},
		tmp:   []int{},
	}
}

func (s *SortedStack) Push(val int) {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, val)
	} else {
		for len(s.stack) > 0 && s.stack[len(s.stack)-1] < val {
			s.tmp = append(s.tmp, s.stack[len(s.stack)-1])
			s.stack = s.stack[:len(s.stack)-1]
		}
		s.stack = append(s.stack, val)
		for len(s.tmp) > 0 {
			s.stack = append(s.stack, s.tmp[len(s.tmp)-1])
			s.tmp = s.tmp[:len(s.tmp)-1]
		}
	}
}

func (s *SortedStack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *SortedStack) Peek() int {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}

	return -1
}

func (s *SortedStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Pop()
	obj.Peek()
	obj.IsEmpty()
}
