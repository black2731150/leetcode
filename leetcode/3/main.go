package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	themap := map[byte]bool{}

	max := 0
	for i := 0; i < len(s); i++ {
		once := 0
		for j := i; j < len(s); j++ {
			if _, isFind := themap[s[j]]; isFind {
				once = 0
				for k := range themap {
					delete(themap, k)
				}
				break
			} else {
				themap[s[j]] = true
				once++
				if once > max {
					max = once
				}
			}
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
