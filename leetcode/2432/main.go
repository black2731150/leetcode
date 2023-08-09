package main

import "fmt"

func hardestWorker(n int, logs [][]int) int {
	time := 0
	max := 0
	index := n
	for i := range logs {
		tmp := logs[i][1] - time
		if tmp >= max {
			if tmp == max {
				max = tmp
				if logs[index][0] > logs[i][0] {
					index = i
				}
			} else {
				max = tmp
				index = i
			}
		}
		time = logs[i][1]
	}

	return logs[index][0]
}

func main() {
	n := 10
	logs := [][]int{{0, 10}, {1, 20}, {3, 30}, {4, 40}}
	fmt.Println(hardestWorker(n, logs))
}
