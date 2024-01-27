package main

import "fmt"

func maximumNumberOfStringPairs(words []string) int {
	theMap := make(map[string]int)
	for _, v := range words {
		theMap[v]++
	}

	answer := 0
	for _, word := range words {
		r := reverse(word)

		if r == word {
			if val, find := theMap[word]; find && val >= 2 {
				theMap[word] -= 2
				answer++
			}
		} else {
			val1, find1 := theMap[word]
			val2, find2 := theMap[r]

			if find1 && find2 && val1 > 0 && val2 > 0 {
				theMap[word]--
				theMap[r]--
				answer++
			}
		}
	}

	return answer
}

func reverse(s string) string {
	arr := []byte(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func main() {
	words := []string{"ab", "ba", "cc"}
	fmt.Println(maximumNumberOfStringPairs(words))
}
