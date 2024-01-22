package main

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	type option struct {
		id     int
		status string
		time   int
	}

	const START = "start"
	const END = "end"

	answer := make([]int, n)

	lastOption := option{}
	stack := []option{}
	for _, log := range logs {
		infers := strings.Split(log, ":")
		id, _ := strconv.Atoi(infers[0])
		status := infers[1]
		time, _ := strconv.Atoi(infers[2])

		if status == START {
			stack = append(stack, option{id, status, time})
			answer[lastOption.id] += time - lastOption.time
		}

		if status == END {
			if stack[len(stack)-1].id == id {
				answer[id] += time - stack[len(stack)-1].time
				stack = stack[:len(stack)-1]
			} else {
				answer[lastOption.id] += time - lastOption.time
			}
		}

		lastOption = option{id, status, time}
	}

	return answer
}
