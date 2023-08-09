package main

func RemoveDuplicates(s string) string {
	if len(s) == 1 {
		return s
	}

	for i := 0; i < len(s)-1; i++ {
		if len(s) == 1 {
			return s
		}

		if i < 0 {
			i = 0
		}

		if s[i] == s[i+1] {
			s = s[:i] + s[i+2:]
			i--
			i--
			continue
		}
	}

	return s
}

// func removeDuplicates(s string) string {
//     stack := []byte{}
//     for i := 0; i < len(s); i++{
//         if i > 0 {
//             if len(stack) > 0 && s[i] == stack[len(stack) - 1]{
//                 stack = stack[0: len(stack) - 1]
//             }else {
//                 stack = append(stack, s[i])
//             }
//         }else{
//             stack = append(stack, s[i])
//         }
//     }
//     return string(stack)
// }
