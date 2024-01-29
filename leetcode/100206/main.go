package main

func maximumLength(nums []int) int {
	numsMap := make(map[int]int)
	oneNUm := 0
	for _, v := range nums {
		if v == 1 {
			oneNUm++
		}
		numsMap[v]++
	}

	answer := 0
	for _, num := range nums {
		tmp := 0
		x := num
		hasone := false
		lastNum := 0
		for {
			n, find := numsMap[x]
			if find {
				if n >= 2 && x != x*x {
					lastNum = n
					tmp += 2
					x = x * x
				} else if n == 1 || (n > 1 && x == x*x) {
					tmp++
					hasone = true
					break
				} else {
					tmp = 1
					break
				}
			} else {
				break
			}
		}
		if !hasone && lastNum > 0 {
			hasone = true
			tmp--
		}

		if !hasone {
			tmp = 1
		}

		if oneNUm%2 == 0 {
			oneNUm--
		}
		answer = max(answer, tmp)
		answer = max(answer, oneNUm)
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{5, 4, 1, 2, 2}
	maximumLength(nums)
}
