package main

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param s string字符串
 * @return string字符串
 */
func decodeString(s string) string {
	stack := []int{}

	sStack := []byte{}

	numStack := []int{}

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			numStack = append(numStack, int(s[i]-'a'))
			continue
		}

		if s[i] == '{' {
			stack = append(stack, i)
			continue
		}

		if s[i] == '}' {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			diff := i - left - 1

			chars := []byte{}
			for j := 0; j < diff; j++ {
				chars = append(chars, sStack[len(sStack)-1])
				sStack = sStack[:len(sStack)-1]
			}

			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			for j := 0; j < num; j++ {
				sStack = append(sStack, chars...)
			}
		}

		if s[i] >= 'a' && s[i] <= 'z' {
			sStack = append(sStack, s[i])
		}
	}

	return string(sStack)
}
