package main

type Bank struct {
	moeny map[int]int64
}

func Constructor(balance []int64) Bank {
	m := make(map[int]int64, len(balance))
	for i, v := range balance {
		m[i+1] = v
	}
	return Bank{
		moeny: m,
	}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	a1, find := b.moeny[account1]
	if !find {
		return false
	}

	if a1-money >= 0 {
		_, find := b.moeny[account2]
		if !find {
			return false
		}
		b.moeny[account1] -= money
		b.moeny[account2] += money
	} else {
		return false
	}

	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	_, find := b.moeny[account]
	if !find {
		return false
	}

	b.moeny[account] += money
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	a, find := b.moeny[account]
	if !find {
		return false
	}
	if a-money >= 0 {
		b.moeny[account] -= money
	} else {
		return false
	}
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */

func main() {
	balance := []int64{10, 100, 20, 50, 30}
	obj := Constructor(balance)
	obj.Transfer(5, 1, 20)
	obj.Deposit(5, 20)
	obj.Withdraw(10, 50)
}
