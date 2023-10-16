package main

func avoidFlood(rains []int) []int {
	n := len(rains)

	canChou := 0

	answer := []int{}

	tmp := []int{}

	theMap := make(map[int]bool, n)
	start := 0

	for i := range rains {
		if rains[i] != 0 {
			start = i
			break
		}
	}

	for i := range rains {
		theMap[rains[i]] = false
	}

	for i, v := range rains {
		if i < start {
			continue
		}

		if v == 0 {
			tmp = append(tmp, i)
			answer = append(answer, 1)
			canChou++
		} else {
			if theMap[v] {
				if canChou > 0 {
					canChou--
					answer[tmp[0]] = v
					tmp = tmp[1:]
				} else {
					return []int{}
				}
			} else {
				theMap[v] = true
			}

			answer = append(answer, -1)
		}
	}
	return answer
}
