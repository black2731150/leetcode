package main

import "fmt"

func beautySum(s string) int {
	n := len(s)
	answer := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			answer += theCha(s[i:j])
		}
	}
	return answer
}

func theCha(s string) int {
	n := len(s)

	themap := make(map[byte]int)
	for i := 0; i < n; i++ {
		themap[s[i]]++
	}

	max := 0
	min := n
	for _, v := range themap {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}

func main() {
	s := "aabcb"
	fmt.Println(beautySum(s))
}
