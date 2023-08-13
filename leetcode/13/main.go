package main

import "fmt"

func romanToInt(s string) int {
	//定义一个map，匹配罗马数字和整数
	RN := make(map[byte]int)
	RN['I'] = 1
	RN['V'] = 5
	RN['X'] = 10
	RN['L'] = 50
	RN['C'] = 100
	RN['D'] = 500
	RN['M'] = 1000

	var answer int
	for i := range s {
		if i < (len(s)-1) && RN[s[i]] < RN[s[i+1]] {
			answer = answer - RN[s[i]]
		} else {
			answer = answer + RN[s[i]]
		}
	}

	return answer
}

func main() {
	s := "LVIII"
	fmt.Println(romanToInt(s))
}
