package main

type SmallestInfiniteSet struct {
	MinNum  int
	addNums map[int]struct{}
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		MinNum:  1,
		addNums: map[int]struct{}{},
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	answer := s.MinNum
	for k := range s.addNums {
		if k < answer {
			answer = k
		}
	}
	if answer == s.MinNum {
		s.MinNum++
	} else {
		delete(s.addNums, answer)
	}

	return answer
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num >= s.MinNum {
		return
	}

	for k := range s.addNums {
		if k == num {
			return
		}
	}

	s.addNums[num] = struct{}{}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

func main() {
	obj := Constructor()
	obj.AddBack(2)
	obj.PopSmallest()
	obj.PopSmallest()
	obj.PopSmallest()
	obj.AddBack(1)
	obj.PopSmallest()
	obj.PopSmallest()
	obj.PopSmallest()
}
