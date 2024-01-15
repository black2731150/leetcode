package main

import (
	"fmt"
	"sort"
)

func beautifulIndices(s string, a string, b string, k int) []int {
	indexi, indexj := make(map[int]struct{}), make(map[int]struct{})
	for i := 0; i < len(s)-len(a)+1; i++ {
		if s[i:i+len(a)] == a {
			indexi[i] = struct{}{}
		}
	}

	for j := 0; j < len(s)-len(b)+1; j++ {
		if s[j:j+len(b)] == b {
			indexj[j] = struct{}{}
		}
	}

	answer := []int{}
	for i := range indexi {
		for j := range indexj {
			if abs(j-i) <= k {
				answer = append(answer, i)
				break
			}
		}
	}

	sort.Ints(answer)
	return answer
}

// 超出时间限制
// func beautifulIndices(s string, a string, b string, k int) []int {
// 	answer := []int{}
// 	for i := 0; i < len(s)-len(a)+1; i++ {
// 		if s[i:i+len(a)] == a {
// 			for j := 0; j < len(s)-len(b)+1; j++ {
// 				if s[j:j+len(b)] == b {
// 					if abs(j-i) <= k {
// 						answer = append(answer, i)
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return answer
// }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	s := "isawsquirrelnearmysquirrelhouseohmy"
	a := "my"
	b := "squirrel"
	k := 15
	fmt.Println(beautifulIndices(s, a, b, k))
}
