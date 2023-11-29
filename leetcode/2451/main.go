package main

import "fmt"

func oddString(words []string) string {
	theMap := make(map[string]int)
	indexMap := make(map[string]int)

	for i, word := range words {
		str := ""
		for j := 0; j < len(word)-1; j++ {
			str = str + string((int(word[j+1] - word[j])))
		}
		theMap[str]++
		indexMap[str] = i

	}

	for k, v := range theMap {
		if v == 1 {
			return words[indexMap[k]]
		}
	}

	return ""
}

func main() {
	words := []string{"adc", "wzy", "abc"}
	fmt.Println(oddString(words))
}
