package main

type MyStack struct {
	InQueue  []int
	OutQueue []int
}

func Constructor() MyStack {
	return MyStack{
		InQueue:  make([]int, 0),
		OutQueue: make([]int, 0),
	}
}

func (ms *MyStack) Push(x int) {

}

func (ms *MyStack) Pop() int {

}

func (ms *MyStack) Top() int {

}

func (ms *MyStack) Empty() bool {
	return len(ms.InQueue) == 0 && len(ms.OutQueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
