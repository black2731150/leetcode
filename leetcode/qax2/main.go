package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param sequence int整型一维数组
 * @return bool布尔型
 */
func CheckGameResource(sequence []int) bool {
	// write code here

	stack := []int{}

	for i := range sequence {
		if sequence[i] == 1 {
			stack = append(stack, 1)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func main() {
	sequence := []int{1, 0, 1, 1, 1, 0, 0, 1, 0, 0}
	fmt.Println(CheckGameResource(sequence))
}
