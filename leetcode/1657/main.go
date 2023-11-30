package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	word1Map, word2Map := make(map[byte]int), make(map[byte]int)

	for i := range word1 {
		word1Map[word1[i]]++
	}

	for i := range word2 {
		word2Map[word2[i]]++
	}

	if len(word1Map) != len(word2Map) {
		return false
	}

	word1Nums, word2Nums := []int{}, []int{}
	for k := range word1Map {
		if word1Map[k] != 0 && word2Map[k] != 0 {

			word1Nums = append(word1Nums, word1Map[k])
			word2Nums = append(word2Nums, word2Map[k])
			continue
		} else {
			return false
		}
	}

	sort.Ints(word1Nums)
	sort.Ints(word2Nums)

	for i, v := range word1Nums {
		if v != word2Nums[i] {
			return false
		}
	}

	return true
}

func main() {
	word1 := "abbzzca"
	word2 := "babzzcz"
	fmt.Println(closeStrings(word1, word2))
}
