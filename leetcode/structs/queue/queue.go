package queue

type Queue struct {
	Nums []int
	Len  int
}

// 构建一个新队列
func NewQueue() *Queue {
	return &Queue{
		Nums: make([]int, 0),
		Len:  0,
	}
}

// 入队一个元素
func (q *Queue) Enqueue(val int) {
	q.Nums = append(q.Nums, val)
	q.Len++
}

// 出队一个元素
func (q *Queue) Dequeue() (int, bool) {
	if q.Len > 0 {
		val := q.Nums[0]
		q.Nums = q.Nums[1:]
		q.Len--
		return val, true
	}

	return 0, false
}

// 查看队首元素
func (q *Queue) Top() (int, bool) {
	if q.Len > 0 {
		return q.Nums[0], true
	}
	return 0, false
}

// 查看队尾元素
func (q *Queue) Last() (int, bool) {
	if q.Len > 0 {
		return q.Nums[q.Len-1], true
	}
	return 0, false
}

// 查看队列是否为空
func (q *Queue) IsEmpty() bool {
	return !(q.Len > 0)
}

// 查看队列的长度
func (q *Queue) Lengh() int {
	return q.Len
}

// 清空队列
func (q *Queue) Clear() {
	q.Nums = q.Nums[:0]
	q.Len = 0
}
