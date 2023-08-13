package main

import "math/rand"

type RandomizedSet struct {
	Set  map[int]int
	Nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Set:  make(map[int]int, 0),
		Nums: make([]int, 0),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.Set[val]; ok {
		return false
	} else {
		r.Set[val] = len(r.Nums)
		r.Nums = append(r.Nums, val)
		return true
	}
}

func (r *RandomizedSet) Remove(val int) bool {
	if index, ok := r.Set[val]; ok {
		r.Nums[index] = r.Nums[len(r.Nums)-1]
		r.Set[r.Nums[index]] = index
		r.Nums = r.Nums[:len(r.Nums)-1]
		delete(r.Set, val)
		return true
	} else {
		return false
	}
}

func (r *RandomizedSet) GetRandom() int {
	return r.Nums[rand.Intn(len(r.Nums))]
}

func main() {
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(2)
	obj.Remove(1)
	obj.GetRandom()
}
