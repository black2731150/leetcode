package main

import "fmt"

func minDeletionSize(strs []string) int {
	lenStr := len(strs)
	lenOnce := len(strs[0])

	once := []byte(strs[0])
	theMap := map[int]bool{}
	for i := 0; i < lenOnce; i++ {
		theMap[i] = false
	}

	answer := 0
	for i := 1; i < lenStr; i++ {
		for j := 0; j < lenOnce; j++ {
			if theMap[j] {
				continue
			}

			if strs[i][j] >= once[j] {
				once[j] = strs[i][j]
			} else {
				answer++
				theMap[j] = true
			}
		}
	}
	return answer
}

func main() {
	strs := []string{"egguij", "gjsnnk", "lstgon", "ztzrqv"}
	fmt.Println(minDeletionSize(strs))
}
