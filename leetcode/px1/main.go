package main

import (
	"fmt"
	"strconv"
)

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 * 返回到0到n之间数字2出现的次数
 * @param n int整型
 * @return int整型
 */
func countOfTwos(n int) int {
	if n == 0 {
		return 0
	}

	sNum := strconv.Itoa(n)

	first := int(sNum[0] - '0')

	elseLen := len(sNum[1:])

	elseNum := 0
	if elseLen > 0 {
		elseNum, _ = strconv.Atoi(sNum[1:])
	}

	var one int

	if first == 2 {
		one = elseNum + 1 + countOfTwos(elseNum)
	}

	if first < 2 {
		one = countOfTwos(elseNum)
	}

	if first > 2 {
		one = (first-1)*countOfTwos(elseNum) + pow10(elseLen) + 1
	}

	return one + countOfTwos(elseNum)
}

func pow10(x int) int {
	answer := 1
	for i := 0; i < x; i++ {
		answer = answer * 10
	}
	return answer
}

func main() {
	fmt.Println(countOfTwos(23))
}
