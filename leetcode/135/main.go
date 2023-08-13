package main

import "fmt"

func candy(ratings []int) int {
	ratingsLen := len(ratings)
	candyNums := make([]int, ratingsLen)

	check := func(x, y int) bool {
		if ratings[x] > ratings[y] {
			if candyNums[x] > candyNums[y] {
				return true
			} else {
				candyNums[x]++
				return false
			}
		} else if ratings[x] == ratings[y] {
			return true
		} else {
			if candyNums[x] < candyNums[y] {
				return true
			} else {
				candyNums[y]++
				return false
			}
		}
	}

	//给每一个孩子发一颗糖
	for i := 0; i < ratingsLen; i++ {
		candyNums[i]++
	}

	//检查相邻两个孩子
	for i := 0; i+1 < ratingsLen; i++ {
		if check(i, i+1) {
			continue
		} else {
			i = 0
		}
	}

	//检查相邻两个孩子
	for i := 0; i+1 < ratingsLen; i++ {
		if check(i, i+1) {
			continue
		} else {
			i = 0
		}
	}

	answer := 0
	for i := range candyNums {
		answer = answer + candyNums[i]
	}

	return answer
}

func main() {
	ratings := []int{3, 2, 1, 1, 4, 3, 3}
	fmt.Println(candy(ratings))
}
