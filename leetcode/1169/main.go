package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Transaction struct {
	name  string
	time  int
	money int
	city  string
	index int
}

func invalidTransactions(transactions []string) []string {
	nameMap := make(map[string][]Transaction)

	answer := []string{}
	for i, transaction := range transactions {
		infos := strings.Split(transaction, ",")
		time, _ := strconv.Atoi(infos[1])
		money, _ := strconv.Atoi(infos[2])

		T := Transaction{
			name:  infos[0],
			time:  time,
			money: money,
			city:  infos[3],
			index: i,
		}

		nameMap[infos[0]] = append(nameMap[infos[0]], T)

	}

	for _, Ts := range nameMap {
		sort.Slice(Ts, func(i, j int) bool {
			return Ts[i].time < Ts[j].time
		})

		for i := 0; i < len(Ts); i++ {
			a := Ts[i]
			if a.money > 1000 {
				answer = append(answer, transactions[a.index])
				continue
			}
			for j := 0; j < len(Ts); j++ {
				if i != j {
					b := Ts[j]
					if b.time-a.time <= 60 && a.city != b.city {
						answer = append(answer, transactions[a.index])
						break
					}
				}

			}
		}
	}
	return answer
}

func main() {
	t := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	fmt.Println(invalidTransactions(t))
}
