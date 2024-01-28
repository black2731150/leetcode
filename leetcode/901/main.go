package main

type StockSpanner struct {
	prices []int
	stack  []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: []int{},
		stack:  []int{},
	}
}

func (s *StockSpanner) Next(price int) int {
	index := len(s.prices)
	s.prices = append(s.prices, price)

	if len(s.stack) == 0 {
		s.stack = append(s.stack, index)
		return 1
	} else {
		for len(s.stack) > 0 && s.prices[s.stack[len(s.stack)-1]] <= price {
			s.stack = s.stack[:len(s.stack)-1]
		}

		if len(s.stack) == 0 {
			s.stack = append(s.stack, index)
			return index + 1
		}

		last := s.stack[len(s.stack)-1]

		s.stack = append(s.stack, index)

		return index - last
	}
}

func main() {
	obj := Constructor()
	obj.Next(31)
	obj.Next(41)
	obj.Next(48)
	obj.Next(59)
	obj.Next(79)
}
