package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func three(woods []int, lenght int) int {
	sort.Ints(woods)
	answer := 0
	for i := len(woods) - 1; i >= 0; {
		if lenght-woods[i] >= 0 {
			lenght = lenght - woods[i]
			answer++
		} else {
			i--
		}
	}
	return answer
}

func main() {
	woodsString := ""
	fmt.Scan(&woodsString)
	lenghtString := ""
	fmt.Scan(&lenghtString)

	fmt.Println(woodsString, lenghtString)

	woodsString = woodsString[1 : len(woodsString)-1]
	strings := strings.Split(woodsString, ",")
	woods := []int{}
	for i := range strings {
		wood, err := strconv.Atoi(strings[i])
		if err != nil {
			break
		}
		woods = append(woods, wood)
	}

	lenght, err := strconv.Atoi(lenghtString)
	if err != nil {
		return
	}

	fmt.Println(three(woods, lenght))
}
