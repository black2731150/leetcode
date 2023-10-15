package main

func distMoney(money int, children int) int {
	money = money - children

	if money < 0 {
		return -1
	}

	answer := 0
	for money > 0 && children > 0 {
		if money >= 7 {
			money = money - 7
			children = children - 1
			answer = answer + 1
		} else {
			if money == 3 {
				if children >= 1 {
					answer = answer - 1
					money = 0
				} else if children == 0 {
					money = 0
				}
			} else {
				money = 0
			}
		}
	}
	return answer
}
