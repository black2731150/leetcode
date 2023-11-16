package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	n := len(s)

	if n == 1 {
		return 1
	}

	chMap := make(map[byte]int)

	maxnum, left := 0, 0
	for i := range s {
		chMap[s[i]] = i
		if len(chMap) > 2 {
			index := n
			for _, v := range chMap {
				if v < index {
					index = v
				}
			}
			delete(chMap, s[index])
			left = index + 1
		}
		maxnum = max(maxnum, i-left+1)
	}

	return maxnum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "abaccc"
	fmt.Println(lengthOfLongestSubstringTwoDistinct(s))
}
