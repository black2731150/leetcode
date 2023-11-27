package queue

type Queue struct {
	Nums []int
}

// 构建一个新队列
func NewQueue() *Queue {
	return &Queue{
		Nums: make([]int, 0),
	}
}

// 入队一个元素
func (q *Queue) Enqueue(val int) {
	q.Nums = append(q.Nums, val)
}

// 出队一个元素
func (q *Queue) Dequeue() (int, bool) {
	if len(q.Nums) > 0 {
		val := q.Nums[0]
		q.Nums = q.Nums[1:]
		return val, true
	}
	return 0, false
}

// 查看队首元素
func (q *Queue) Top() (int, bool) {
	if len(q.Nums) > 0 {
		return q.Nums[0], true
	}
	return 0, false
}

// 查看队尾元素
func (q *Queue) Last() (int, bool) {
	if len(q.Nums) > 0 {
		return q.Nums[len(q.Nums)-1], true
	}
	return 0, false
}

// 查看队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.Nums) == 0
}

// 查看队列的长度
func (q *Queue) Length() int {
	return len(q.Nums)
}

// 清空队列
func (q *Queue) Clear() {
	q.Nums = q.Nums[:0]
}
