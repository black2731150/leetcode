package main

import "fmt"

func findLongestWord(s string, dictionary []string) string {
	n := len(s)
	var answer string
	for _, str := range dictionary {
		sIndex, strIndex := 0, 0
		for strIndex < len(str) && sIndex < n {
			if str[strIndex] == s[sIndex] {
				strIndex++
			}
			sIndex++
		}

		if strIndex == len(str) { // 只有当 strIndex 遍历完整个 str 时，才表示 str 可以由 s 删除某些字符得到
			if len(str) > len(answer) || (len(str) == len(answer) && str < answer) {
				answer = str
			}
		}
	}

	return answer
}

func main() {
	s := "abcd"
	dictionary := []string{"a", "ab", "abc"}
	fmt.Println(findLongestWord(s, dictionary))
}

// func findLongestWord(s string, dictionary []string) string {
// 	n := len(s)
// 	var answer string
// 	for _, str := range dictionary {
// 		sIndex := 0
// 		for i := 0; i < len(str); {
// 			if str[i] == s[sIndex] {
// 				i++
// 			}

// 			sIndex++

// 			if sIndex == n {
// 				break
// 			}
// 		}

// 		if sIndex < n {
// 			if len(answer) == len(str) {
// 				x := []string{str, answer}
// 				sort.Strings(x)
// 				answer = x[0]
// 			} else if len(answer) < len(str) {
// 				answer = str
// 			}
// 		}
// 	}

// 	return answer
// }
