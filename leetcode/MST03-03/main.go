package main

type StackOfPlates struct {
	stackNum int
	stacks   [][]int
	cap      int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		stackNum: 0,
		stacks:   [][]int{},
		cap:      cap,
	}
}

func (s *StackOfPlates) Push(val int) {
	if s.cap <= 0 {
		return
	}

	if s.stackNum == 0 || len(s.stacks[s.stackNum-1]) == s.cap {
		s.stacks = append(s.stacks, []int{val})
		s.stackNum++
	} else {
		s.stacks[s.stackNum-1] = append(s.stacks[s.stackNum-1], val)
	}
}

func (s *StackOfPlates) Pop() int {
	if s.stackNum == 0 {
		return -1
	}

	index := s.stackNum - 1
	value := s.stacks[index][len(s.stacks[index])-1]
	s.stacks[index] = s.stacks[index][:len(s.stacks[index])-1]

	if len(s.stacks[index]) == 0 {
		s.stacks = s.stacks[:index]
		s.stackNum--
	}

	return value
}

func (s *StackOfPlates) PopAt(index int) int {
	if index < 0 || index >= s.stackNum || len(s.stacks[index]) == 0 {
		return -1
	}

	value := s.stacks[index][len(s.stacks[index])-1]
	s.stacks[index] = s.stacks[index][:len(s.stacks[index])-1]

	if len(s.stacks[index]) == 0 {

		if index != s.stackNum-1 {
			s.stacks = append(s.stacks[:index], s.stacks[index+1:]...)
		} else {
			s.stacks = s.stacks[:index]
		}
		s.stackNum--
	}

	return value
}

func main() {

	obj := Constructor(2)
	obj.Push(1)
	obj.Pop()
	obj.PopAt(0)

}
