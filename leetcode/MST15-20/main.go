package main

import "fmt"

func getValidT9Words(num string, words []string) []string {
	theMap := map[byte]string{}

	theMap['2'] = "abc"
	theMap['3'] = "def"
	theMap['4'] = "ghi"
	theMap['5'] = "jkl"
	theMap['6'] = "mno"
	theMap['7'] = "pqrs"
	theMap['8'] = "tuv"
	theMap['9'] = "wxyz"

	answer := []string{}
	for _, word := range words {

		j := 0
		for ; j < len(num); j++ {
			letters := theMap[num[j]]
			isFind := false
			for k := 0; k < len(letters); k++ {
				if letters[k] == word[j] {
					isFind = true
					break
				}
			}
			if isFind {
				continue
			} else {
				break
			}
		}

		if j == len(num) {
			answer = append(answer, word)
		}
	}

	return answer
}

func main() {
	num := "8733"
	words := []string{"tree", "used"}
	fmt.Println(getValidT9Words(num, words))
}
