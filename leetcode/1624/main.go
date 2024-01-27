package main

import "fmt"

func maxLengthBetweenEqualCharacters(s string) int {
	chMap := make(map[byte][]int)
	for i := range s {
		ch := s[i]
		if _, find := chMap[ch]; find {
			chMap[ch][1] = i
		} else {
			chMap[ch] = []int{i, -1}
		}
	}

	answer := -1
	for _, arr := range chMap {
		minIndex := arr[0]
		maxIndex := arr[1]
		answer = max(maxIndex-minIndex-1, answer)
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "ansijninadsd"
	fmt.Println(maxLengthBetweenEqualCharacters(s))
}
