package main

func isValid(s string) bool {
	lenS := len(s)
	if lenS%2 != 0 {
		return false
	}
	stack := make([]byte, lenS)
	idnex := -1

	for i := range s {
		if isLeft(s[i]) {
			idnex++
			stack[idnex] = s[i]
		} else {

			if idnex < 0 {
				return false
			}

			if s[i] == ')' {
				if stack[idnex] == '(' {
					idnex--
					continue
				} else {
					return false
				}
			}

			if s[i] == ']' {
				if stack[idnex] == '[' {
					idnex--
					continue
				} else {
					return false
				}
			}

			if s[i] == '}' {
				if stack[idnex] == '{' {
					idnex--
					continue
				} else {
					return false
				}
			}
		}
	}

	return idnex == -1
}

func isLeft(ch byte) bool {
	if ch == '(' || ch == '[' || ch == '{' {
		return true
	}
	return false
}
