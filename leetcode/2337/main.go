package main

func canChange(start string, target string) bool {
	sindex := 0
	for i := 0; i < len(target); i++ {
		if target[i] == '_' {
			continue
		} else {
			if target[i] == 'L' {
				for ; sindex < len(start); sindex++ {
					if start[sindex] == '_' {
						continue
					} else {
						if start[sindex] == 'L' {
							sindex++
							break
						} else {
							return false
						}
					}
				}
			} else {
				for ; sindex < len(start); sindex++ {
					if start[sindex] == '_' {
						continue
					} else {
						if start[sindex] == 'R' {
							sindex++
							break
						} else {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
