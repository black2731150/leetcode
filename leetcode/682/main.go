package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	var answerArr []int = []int{}
	for _, v := range operations {
		if v == "+" || v == "D" || v == "C" {
			if v == "+" {
				num := answerArr[len(answerArr)-1] + answerArr[len(answerArr)-2]
				answerArr = append(answerArr, num)
			}

			if v == "D" {
				num := answerArr[len(answerArr)-1] * 2
				answerArr = append(answerArr, num)
			}

			if v == "C" {
				answerArr = answerArr[:len(answerArr)-1]
			}
		} else {
			num, _ := strconv.Atoi(v)
			answerArr = append(answerArr, num)
		}
	}

	answer := 0
	for _, v := range answerArr {
		answer = answer + v
	}
	return answer
}

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
}
