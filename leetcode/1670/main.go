package main

type FrontMiddleBackQueue struct {
	nums []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		nums: []int{},
	}
}

func (f *FrontMiddleBackQueue) PushFront(val int) {
	if len(f.nums) == 0 {
		f.nums = append(f.nums, val)
	} else {
		f.nums = append([]int{val}, f.nums...)
	}
}

func (f *FrontMiddleBackQueue) PushMiddle(val int) {

	if len(f.nums) == 0 {
		f.nums = append(f.nums, val)
		return
	}

	mid := len(f.nums) / 2
	left := f.nums[:mid]
	right := f.nums[mid:]

	answer := []int{}
	answer = append(answer, left...)
	answer = append(answer, val)
	answer = append(answer, right...)
	f.nums = answer
}

func (f *FrontMiddleBackQueue) PushBack(val int) {
	f.nums = append(f.nums, val)
}

func (f *FrontMiddleBackQueue) PopFront() int {
	if len(f.nums) == 0 {
		return -1
	}
	answer := f.nums[0]
	f.nums = f.nums[1:]
	return answer
}

func (f *FrontMiddleBackQueue) PopMiddle() int {
	if len(f.nums) == 0 {
		return -1
	}

	mid := (len(f.nums) - 1) / 2
	answer := f.nums[mid]
	f.nums = append(f.nums[:mid], f.nums[mid+1:]...)
	return answer
}

func (f *FrontMiddleBackQueue) PopBack() int {
	if len(f.nums) == 0 {
		return -1
	}

	answer := f.nums[len(f.nums)-1]
	f.nums = f.nums[:len(f.nums)-1]
	return answer
}

func main() {
	obj := Constructor()
	obj.PushFront(1)
	obj.PushMiddle(2)
	obj.PushBack(3)
	obj.PopFront()
	obj.PopMiddle()
	obj.PopBack()
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
