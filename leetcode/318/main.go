package main

import (
	"fmt"
	"sort"
	"strings"
)

// func maxProduct(words []string) int {
// 	answer := 0

// 	for i := range words {
// 		words[i] = sortString(words[i])
// 	}

// 	sort.Strings(words)

// 	for i := 0; i < len(words); i++ {
// 		for j := i + 1; j < len(words); j++ {
// 			if words[i][0] == words[j][0] {
// 				continue
// 			}

// 			if !publicCh(words[i], words[j]) {
// 				if x := len(words[i]) * len(words[j]); answer < x {
// 					answer = x
// 				}
// 			}
// 		}

// 	}
// 	return answer
// }

func publicCh(s1, s2 string) bool {
	themap := make(map[rune]int)
	for _, v := range s1 {
		themap[v]++
	}

	for _, v := range s2 {
		if _, ok := themap[v]; ok {
			return true
		}
	}

	return false
}

func sortString(s string) string {
	// 将字符串转换为字符数组
	chars := strings.Split(s, "")

	// 对字符数组进行排序
	sort.Strings(chars)

	// 将排序后的字符数组重新组合成字符串
	return strings.Join(chars, "")
}

func main() {
	words := []string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"}
	fmt.Println(maxProduct(words))
}

func maxProduct(words []string) int {
	wordMap := make(map[string]int)
	for _, word := range words {
		val := 0
		for _, ch := range word {
			val |= 1 << (ch - 'a')
		}
		wordMap[word] = val
	}

	answer := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if wordMap[words[i]]&wordMap[words[j]] == 0 {
				if x := len(words[i]) * len(words[j]); answer < x {
					answer = x
				}
			}
		}
	}
	return answer
}
