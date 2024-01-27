package main

func intToRoman(num int) string {
	theMap := make(map[int]byte, 7)
	theMap[1] = 'I'
	theMap[5] = 'V'
	theMap[10] = 'X'
	theMap[50] = 'L'
	theMap[100] = 'C'
	theMap[500] = 'D'
	theMap[1000] = 'M'

	answer := []byte{}
	for num > 0 {
		if num >= 1000 {
			answer = append(answer, theMap[1000])
			num = num - 1000
			continue
		}

		if num > 800 {
			for num > 800 {
				answer = append(answer, theMap[100])
				num -= 100
			}
			answer = append(answer, theMap[])
		}

		if num >= 500 {
			answer = append(answer, theMap[500])
			num -= 500
			continue
		}

		if num >= 100 {
			answer = append(answer, theMap[100])
			num -= 100
			continue
		}

		if num >= 50 {
			answer = append(answer, theMap[50])
			num -= 50
			continue
		}

		if num >= 10 {
			answer = append(answer, theMap[10])
			num -= 10
			continue
		}

		if num >= 5 {
			answer = append(answer, theMap[5])
			num -= 5
			continue
		}

		if num >= 1 {
			answer = append(answer, theMap[1])
			num -= 1
			continue
		}
	}

	return string(answer)
}
