package main

import "fmt"

func main() {
	nums := make([]int, 13)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	answer := f(nums)
	if len(answer) == 0 {
		fmt.Println(0)
	} else {
		for _, v := range answer {
			fmt.Print(v, " ")
		}
	}

}

func f(nums []int) []int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}
	answer := []int{}
	for k, v := range numMap {
		if v == 1 {
			cm := make(map[int]int, len(numMap))
			for k2, v2 := range numMap {
				cm[k2] = v2
			}
			delete(cm, k)
			_, can := shunzi(cm, 0)
			if can {
				answer = append(answer, k)
			}

			cm = make(map[int]int, len(numMap))
			for k2, v2 := range numMap {
				cm[k2] = v2
			}
			delete(cm, k)
			_, can = kezi(cm, 0)
			if can {
				answer = append(answer, k)
			}
			continue
		}

		if v >= 2 {
			cm := make(map[int]int, len(numMap))
			for k2, v2 := range numMap {
				cm[k2] = v2
			}
			cm[k] -= 2
			ans, can := shunzi(cm, 1)
			if can {
				answer = append(answer, ans)
			}

			cm = make(map[int]int, len(numMap))
			for k2, v2 := range numMap {
				cm[k2] = v2
			}
			cm[k] -= 2
			ans, can = kezi(cm, 1)
			if can {
				answer = append(answer, ans)
			}
			continue
		}
	}

	return answer
}

func shunzi(theMap map[int]int, cuo int) (int, bool) {
	ans := -1
	can := true
	for x := 0; x < 4; x++ {
		mink := 100
		for k := range theMap {
			mink = min(mink, k)
		}

		if mink == 8 {
			if cuo > 0 {
				cuo--
				ans = 7
				return ans, can
			}
		}

		for i := 0; i < 3; i++ {
			if n, find := theMap[mink+i]; find && n > 0 {
				theMap[mink+i]--
				if theMap[mink+i] == 0 {
					delete(theMap, mink+i)
				}
			} else {
				if cuo > 0 {
					cuo--
					ans = mink + i
				} else {
					can = false
					return ans, can
				}
			}
		}
	}

	return ans, can
}

func kezi(theMap map[int]int, cuo int) (int, bool) {
	ans := -1
	can := true
	for k, v := range theMap {
		if v != 3 {
			if cuo > 0 {
				cuo--
				ans = k
			} else {
				can = false
				return ans, can
			}
		}
	}
	return ans, can
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
