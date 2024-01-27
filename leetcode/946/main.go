package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	n := len(popped)

	stack := []int{}
	pushIndex := 0
	popIndex := 0

	for pushIndex < n {
		if len(stack) == 0 {
			stack = append(stack, pushed[pushIndex])
			pushIndex++
		} else {
			for pushIndex < n && stack[len(stack)-1] != popped[popIndex] {
				stack = append(stack, pushed[pushIndex])
				pushIndex++
			}
			for len(stack) > 0 && stack[len(stack)-1] == popped[popIndex] {
				stack = stack[:len(stack)-1]
				popIndex++
			}
		}
	}

	if pushIndex == popIndex {
		return true
	} else {
		for len(stack) > 0 && popIndex < n {
			if stack[len(stack)-1] == popped[popIndex] {
				popIndex++
				continue
			} else {
				return false
			}
		}
	}

	return pushIndex == popIndex
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	poped := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, poped))
}
