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
			for i := 0; i < num/1000; i++ {
				answer = append(answer, theMap[1000])
				num = num - 1000
			}
		}

		if num > 800 {
			answer = append(answer, theMap[1000])

		}
	}
}
