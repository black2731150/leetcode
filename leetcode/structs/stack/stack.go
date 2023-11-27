package stack

type Stack struct {
	Nums []int
}

// 新建一个栈结构
func NewStack() *Stack {
	return &Stack{
		Nums: make([]int, 0),
	}
}

// 入栈一个元素
func (s *Stack) Push(val int) {
	s.Nums = append(s.Nums, val)
}

// 出栈一个元素
func (s *Stack) Pop() (int, bool) {
	if len(s.Nums) > 0 {
		val := s.Nums[len(s.Nums)-1]
		s.Nums = s.Nums[:len(s.Nums)-1]
		return val, true
	}
	return 0, false
}

// 查看栈顶元素
func (s *Stack) Top() (int, bool) {
	if len(s.Nums) > 0 {
		return s.Nums[len(s.Nums)-1], true
	}
	return 0, false
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.Nums) == 0
}

// 获取栈的长度
func (s *Stack) Length() int {
	return len(s.Nums)
}

// 清空栈
func (s *Stack) Clear() {
	s.Nums = s.Nums[:0]
}

// 获取整个栈
func (s *Stack) Range() []int {
	return s.Nums
}
