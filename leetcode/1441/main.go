package main

import "fmt"

func buildArray(target []int, n int) []string {
	answer := []string{}

	stack := make([]int, n)
	for i := n; i >= 1; i-- {
		stack[i-1] = n - i + 1
	}

	for _, v := range target {
		if stack[len(stack)-1] == v {
			answer = append(answer, "Push")
			stack = stack[:len(stack)-1]
		} else {
			pushNum := 0
			for stack[len(stack)-1] != v {
				stack = stack[:len(stack)-1]
				answer = append(answer, "Push")
				pushNum++
			}

			for i := 0; i < pushNum; i++ {
				answer = append(answer, "Pop")
			}

			if stack[len(stack)-1] == v {
				answer = append(answer, "Push")
				stack = stack[:len(stack)-1]
			}
		}
	}
	return answer
}

func main() {
	target := []int{1, 3}
	n := 3
	fmt.Println(buildArray(target, n))
}
