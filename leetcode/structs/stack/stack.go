package stack

type Stack struct {
	Nums []int
	Len  int
}

// 新建一个栈结构
func NewStack() *Stack {
	return &Stack{
		Nums: make([]int, 0),
		Len:  0,
	}
}

// 入栈一个元素
func (s *Stack) Push(val int) {
	s.Nums = append(s.Nums, val)
	s.Len++
}

// 出栈一个元素
func (s *Stack) Pop() (int, bool) {
	if s.Len > 0 {
		val := s.Nums[len(s.Nums)-1]
		s.Nums = s.Nums[:s.Len-1]
		s.Len--
		return val, true
	}

	return 0, false
}

// 查看栈顶元素
func (s *Stack) Top() (int, bool) {
	if s.Len > 0 {
		val := s.Nums[len(s.Nums)-1]
		return val, true
	}

	return 0, false
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return !(s.Len > 0)
}

// 获取栈的长度
func (s *Stack) Lenght() int {
	return s.Len
}

// 清空栈
func (s *Stack) Clear() {
	s.Nums = s.Nums[:0]
	s.Len = 0
}

// 获取整个栈
func (s *Stack) Range() []int {
	return s.Nums
}
