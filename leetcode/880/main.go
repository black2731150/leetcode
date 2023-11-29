package main

import "fmt"

func decodeAtIndex(s string, k int) string {
	length := 0

	// 第一步：计算解码字符串的长度
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			length *= int(s[i] - '0')
		} else {
			length++
		}
	}

	// 第二步：逆向查找第 k 个字符
	for i := len(s) - 1; i >= 0; i-- {
		k %= length
		if k == 0 && !(s[i] >= '0' && s[i] <= '9') {
			return string(s[i])
		}

		if s[i] >= '0' && s[i] <= '9' {
			length /= int(s[i] - '0')
		} else {
			length--
		}
	}

	return ""
}

//超出内存限制
// func decodeAtIndex(s string, k int) string {
// 	answer := ""

// 	for i := range s {

// 		if len(answer) >= k {
// 			fmt.Println(answer)
// 			return string(answer[k-1])
// 		}

// 		if s[i] >= '0' && s[i] <= '9' {
// 			d := int(s[i] - '0')
// 			x := answer
// 			for i := 1; i < d; i++ {
// 				answer = answer + x
// 			}
// 		} else {
// 			answer = answer + string(s[i])
// 		}
// 	}

// 	return string(answer[k-1])
// }

func main() {
	s := "leet2code3"
	k := 10
	fmt.Println(decodeAtIndex(s, k))
}
