package main

import "fmt"

func countSeniors(details []string) int {
	answer := 0
	for _, v := range details {
		if v[11] >= '6' {
			if v[11] == '6' {
				if v[12] > '0' {
					answer++
				}
			} else {
				answer++
			}
		}

	}
	return answer
}

func main() {
	details := []string{"5612624052M0130", "5378802576M6424", "5447619845F0171", "2941701174O9078"}
	fmt.Println(countSeniors(details))
}
