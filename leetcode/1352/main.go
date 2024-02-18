package main

type ProductOfNumbers struct {
	nums []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		nums: []int{},
	}
}

func (p *ProductOfNumbers) Add(num int) {
	p.nums = append(p.nums, num)
	for i := len(p.nums) - 2; i >= 0; i-- {
		p.nums[i] = p.nums[i] * num
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	return p.nums[len(p.nums)-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

func main() {
	p := Constructor()
	p.Add(1)
	p.Add(2)
	p.Add(3)
	p.GetProduct(2)
}
