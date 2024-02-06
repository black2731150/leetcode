package main

func findLongestSubarray(array []string) []string {
	answer := []int{0, 0}

	type data struct {
		letters int
		nums    int
	}

	qianzhui := make([]data, len(array)+1)

	theMap := make(map[int]int)
	theMap[0] = 0

	for i := 0; i < len(array); i++ {
		if IsLeeter(array[i]) {
			last := qianzhui[i]
			last.letters++
			qianzhui[i+1] = last

			if indexs, find := theMap[last.letters-last.nums]; find {
				if i-indexs > answer[1]-answer[0] {
					answer[1] = i
					answer[0] = indexs
				}
			} else {
				theMap[last.letters-last.nums] = i
			}
		} else {
			last := qianzhui[i]
			last.nums++
			qianzhui[i+1] = last
			if indexs, find := theMap[last.letters-last.nums]; find {
				if i-indexs > answer[1]-answer[0] {
					answer[1] = i
					answer[0] = indexs
				}
			} else {
				theMap[last.letters-last.nums] = i
			}
		}
	}

	return array[answer[0]:answer[1]]
}

func IsLeeter(s string) bool {
	ch := s[0]
	if ch >= 'A' && ch <= 'Z' {
		return true
	}

	return false
}
