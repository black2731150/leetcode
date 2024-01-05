package main

func removeDuplicateLetters(s string) string {
	bytes := make([]bool, 26)
	for i := range s {
		bytes[s[i]-'a'] = true
	}

	answer := []byte{}
	for i, v := range bytes {
		if v {
			answer = append(answer, byte('a'+i))
		}
	}
	return string(answer)
}
