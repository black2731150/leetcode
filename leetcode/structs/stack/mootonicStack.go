package stack

//单调递增栈
type MonotonicIncrementStack struct {
	Nums []int
}

//单调递减栈
type MonotonicDecreaseStack struct {
	Nums []int
}

//创建一个单调递增栈
func NewMonotonicIncrementStack() *MonotonicIncrementStack {
	return &MonotonicIncrementStack{
		Nums: []int{},
	}
}

//创建一个单调递减栈
func NewMonotonicDecreasStack() *MonotonicDecreaseStack {
	return &MonotonicDecreaseStack{
		Nums: []int{},
	}
}

//给单调递增栈入栈一个元素
func (m *MonotonicIncrementStack) Push(val int) []int {
	outVals := []int{}
	if len(m.Nums) == 0 {
		m.Nums = append(m.Nums, val)
	} else {
		for len(m.Nums) > 0 && m.Nums[len(m.Nums)-1] < val {
			outVals = append(outVals, m.Nums[len(m.Nums)-1])
			m.Nums = m.Nums[:len(m.Nums)-1]
		}
		m.Nums = append(m.Nums, val)
	}
	return outVals
}

//从单调递增栈出栈一个元素
func (m *MonotonicIncrementStack) Pop() (int, bool) {
	if len(m.Nums) == 0 {
		return 0, false
	} else {
		answer := m.Nums[len(m.Nums)-1]
		m.Nums = m.Nums[:len(m.Nums)-1]
		return answer, true
	}
}

//给单调递减栈入栈一个元素
func (m *MonotonicDecreaseStack) Push(val int) []int {
	outVals := []int{}
	if len(m.Nums) == 0 {
		m.Nums = append(m.Nums, val)
	} else {
		for len(m.Nums) > 0 && m.Nums[len(m.Nums)-1] > val {
			outVals = append(outVals, m.Nums[len(m.Nums)-1])
			m.Nums = m.Nums[:len(m.Nums)-1]
		}
		m.Nums = append(m.Nums, val)
	}
	return outVals
}

//从单调递减栈出栈一个元素
func (m *MonotonicDecreaseStack) Pop() (int, bool) {
	if len(m.Nums) == 0 {
		return 0, false
	} else {
		answer := m.Nums[len(m.Nums)-1]
		m.Nums = m.Nums[:len(m.Nums)-1]
		return answer, true
	}
}

// Length 方法返回栈的长度
func (s *MonotonicIncrementStack) Length() int {
	return len(s.Nums)
}

// Top 方法返回栈顶元素
func (s *MonotonicIncrementStack) Top() (int, bool) {
	if len(s.Nums) > 0 {
		return s.Nums[len(s.Nums)-1], true
	}
	return 0, false
}

// IsEmpty 方法检查栈是否为空
func (s *MonotonicIncrementStack) IsEmpty() bool {
	return len(s.Nums) == 0
}

// Clear 方法清空栈
func (s *MonotonicIncrementStack) Clear() {
	s.Nums = []int{}
}

// Range 方法返回栈的全部元素
func (s *MonotonicIncrementStack) Range() []int {
	return s.Nums
}

// Length 方法返回栈的长度
func (s *MonotonicDecreaseStack) Length() int {
	return len(s.Nums)
}

// Top 方法返回栈顶元素
func (s *MonotonicDecreaseStack) Top() (int, bool) {
	if len(s.Nums) > 0 {
		return s.Nums[len(s.Nums)-1], true
	}
	return 0, false
}

// IsEmpty 方法检查栈是否为空
func (s *MonotonicDecreaseStack) IsEmpty() bool {
	return len(s.Nums) == 0
}

// Clear 方法清空栈
func (s *MonotonicDecreaseStack) Clear() {
	s.Nums = []int{}
}

// Range 方法返回栈的全部元素
func (s *MonotonicDecreaseStack) Range() []int {
	return s.Nums
}
