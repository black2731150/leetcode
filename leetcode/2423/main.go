package main

import "fmt"

func equalFrequency(word string) bool {
	theMap := make(map[byte]int)
	for i := range word {
		theMap[word[i]]++
	}

	for ch, chNum := range theMap {
		if chNum == 1 {
			num := 0
			for c, v := range theMap {
				if v > 0 && c != ch {
					num = v
					break
				}
			}

			ok := true
			for c, v := range theMap {
				if c != ch {
					if v == num {
						continue
					} else {
						ok = false
						break
					}
				}
			}

			if ok {
				return true
			} else {
				continue
			}
		} else {
			num := chNum - 1
			ok := true
			for c, v := range theMap {
				if c != ch {
					if v == num {
						continue
					} else {
						ok = false
						break
					}
				}
			}
			if ok {
				return true
			} else {
				continue
			}
		}
	}
	return false
}

func main() {
	word := "abcc"
	fmt.Println(equalFrequency(word))
}
