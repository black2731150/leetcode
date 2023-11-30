package main

import (
	"fmt"
	"sort"
)

type CharNums struct {
	char byte
	num  int
}

func minimumKeypresses(s string) int {
	charNums := make([]CharNums, 26)
	for _, ch := range s {
		charNums[ch-'a'].char = byte(ch)
		charNums[ch-'a'].num++
	}

	sort.Slice(charNums, func(i, j int) bool {
		return charNums[i].num > charNums[j].num
	})

	answer := 0

	stack1, stack2, stack3 := make(map[byte]int), make(map[byte]int), make(map[byte]int)
	for _, cn := range charNums {
		if len(stack1) < 9 {
			stack1[cn.char] = 1
		} else {
			if len(stack2) < 9 {
				stack2[cn.char] = 2
			} else {
				stack3[cn.char] = 3
			}
		}
	}

	theMap := make(map[byte]int, 26)

	for k, v := range stack1 {
		theMap[k] = v
	}

	for k, v := range stack2 {
		theMap[k] = v
	}

	for k, v := range stack3 {
		theMap[k] = v
	}

	for i := range s {
		ch := s[i]
		answer += theMap[ch]
	}

	return answer
}

func main() {
	s := "doshdsandosamdasndoad"
	fmt.Println(minimumKeypresses(s))
}
