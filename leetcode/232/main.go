package main

type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{
		input:  []int{},
		output: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	if len(this.output) == 0 {
		this.output = append(this.output, this.input...)
		this.input = []int{}
	}
	answer := this.output[0]
	if len(this.output) >= 2 {
		this.output = this.output[1:]
	} else {
		this.output = []int{}
	}
	return answer
}

func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		this.output = append(this.output, this.input...)
		this.input = []int{}
	}
	return this.output[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.output) == 0 {
		if len(this.input) == 0 {
			return true
		}
	}
	return false
}
