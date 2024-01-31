package main

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	n1 := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		n1[int(s1[i]-'a')]++
		n1[int(s2[i]-'a')]--
	}

	for _, v := range n1 {
		if v != 0 {
			return false
		}
	}

	return true
}
