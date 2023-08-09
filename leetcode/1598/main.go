package main

import "fmt"

func minOperations(logs []string) int {
	answer := 0
	for i := range logs {
		if logs[i] == "../" {
			if answer != 0 {
				answer--
			}
			continue
		}

		if logs[i] == "./" {
			continue
		}

		answer++
	}
	return answer
}

func main() {
	logs := []string{"d1/", "../", "../", "../"}
	fmt.Println(minOperations(logs))
}
