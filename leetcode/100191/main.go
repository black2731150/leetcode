package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	arr := make([]int, 26)
	for i := range word {
		arr[word[i]-'a']++
	}

	sort.Ints(arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	answer := 0

	onenum, townum, threenum := 0, 0, 0
	for i := 0; i < 26; i++ {
		if arr[i] == 0 {
			continue
		} else {
			if onenum < 8 {
				onenum++
				answer += arr[i] * 1
			} else {
				if townum < 8 {
					townum++
					answer += arr[i] * 2
				} else {
					if threenum < 8 {
						threenum++
						answer += arr[i] * 3
					} else {
						answer += arr[i] * 4
					}
				}
			}
		}
	}

	return answer

}

func main() {
	word := "aabbccddeeffgghhiiiiii"
	fmt.Println(minimumPushes(word))
}
