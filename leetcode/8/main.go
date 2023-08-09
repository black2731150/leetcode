package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {

	if len(s) == 0 {
		return 0
	}

	// 读入字符串并丢弃无用的前导空格
	var begin int
	for be, letter := range s {
		if letter == ' ' {
			continue
		} else {
			begin = be
			break
		}
	}

	// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
	var sign byte = '+'
	if s[begin] == '-' {
		sign = '-'
		begin++
	} else if s[begin] == '+' {
		sign = '+'
		begin++
	}

	// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
	// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
	var answer int32
	for _, num := range s[begin:] {
		if isNum(num) {
			temp := answer
			answer = answer*10 + int32((num)) - '0'
			if answer/10 != temp {
				if sign == '+' {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		} else {
			break
		}
	}

	// 如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。

	// 返回整数作为最终结果。
	if sign == '-' {
		return -int(answer)
	} else {
		return int(answer)
	}

}

func isNum(num rune) bool {
	if num >= '0' && num <= '9' {
		return true
	}
	return false
}

func main() {
	num := "12345"
	fmt.Println(myAtoi(num))
}
