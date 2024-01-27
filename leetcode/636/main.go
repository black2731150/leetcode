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

		if 

		lastOption = option{id, status, time}
	}

	return answer
}
