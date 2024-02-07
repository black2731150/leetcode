package main

import "fmt"

func masterMind(solution string, guess string) []int {
	a1, a2, a3, a4 := solution[0], solution[1], solution[2], solution[3]
	b1, b2, b3, b4 := guess[0], guess[1], guess[2], guess[3]

	aMap, bMap := make(map[byte]int), make(map[byte]int)

	caizhong := 0
	if a1 != b1 {
		aMap[a1]++
		bMap[b1]++
	} else {
		caizhong++
	}
	if a2 != b2 {
		aMap[a2]++
		bMap[b2]++
	} else {
		caizhong++
	}
	if a3 != b3 {
		aMap[a3]++
		bMap[b3]++
	} else {
		caizhong++
	}
	if a4 != b4 {
		aMap[a4]++
		bMap[b4]++
	} else {
		caizhong++
	}

	weiCaiZhong := 0
	for k, v := range bMap {
		if n, find := aMap[k]; find {
			for n > 0 && v > 0 {
				weiCaiZhong++
				aMap[k]--
				n--
				v--
				bMap[k]--
			}
		}
	}

	return []int{caizhong, weiCaiZhong}
}

func main() {
	solution := "RGRB"
	guess := "BBBY"
	fmt.Println(masterMind(solution, guess))
}
