package main

func calculate(s string) int {
	numStack, sStack := []int{}, []byte{}
	for i := range s {
		if s[i] >= 0 && s[i] <= 9 {
			numStack = append(numStack, int(s[i]-'0'))
			if len(sStack) > 0 && (sStack[len(sStack)-1] == '*' || sStack[len(sStack)-1] == '/') {
				a := numStack[len(numStack)-1]
				b := numStack[len(numStack)-2]
				if sStack[len(sStack)-1] == '*' {
					a = b * a
				} else {
					a = b / a
				}
				sStack = sStack[:len(sStack)-1]

				numStack = numStack[:len(numStack)-2]
				numStack = append(numStack, a)
			}
		} else {
			sStack = append(sStack, s[i])
		}
	}

	return numStack[0]
}
