package main

import "fmt"

type TwoSum struct {
	nums map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		nums: make(map[int]int),
	}
}

func (t *TwoSum) Add(number int) {
	t.nums[number]++
}

func (t *TwoSum) Find(value int) bool {
	for num := range t.nums {
		x := value - num
		if x == num {
			if count, ok := t.nums[x]; ok && count >= 2 {
				return true
			}
		} else {
			if _, ok := t.nums[x]; ok {
				return true
			}
		}
	}
	return false
}

func main() {
	ovj := Constructor()
	ovj.Add(1)
	ovj.Add(2)
	fmt.Println(ovj.Find(3))
}
