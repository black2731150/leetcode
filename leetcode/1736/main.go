package main

import "fmt"

func maximumTime(time string) string {
	first, second, thred, fourth := time[0], time[1], time[3], time[4]

	answer := ""

	if first == '?' {
		if second != '?' {
			if second > '3' {
				first = '1'
			} else {
				first = '2'
			}
		} else {
			first = '2'
		}

	}

	if second == '?' {
		if first == '2' {
			second = '3'
		} else {
			second = '9'
		}
	}

	if thred == '?' {
		thred = '5'
	}

	if fourth == '?' {
		fourth = '9'
	}

	answer = string(first) + string(second) + ":" + string(thred) + string(fourth)
	return answer
}

func main() {
	time := "0?:3?"
	fmt.Println(maximumTime(time))
}
