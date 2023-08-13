package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		j := 1
		start := i
		if gas[i] >= cost[i] {

			allGas := gas[i]
			allGas = allGas - cost[i]

			for allGas >= 0 {
				allGas = allGas + gas[(start+j)%len(gas)]
				allGas = allGas - cost[(start+j)%len(cost)]

				if (start+j)%len(gas) == start {
					return start
				}
				j++
			}
		}
		i = (start + j + 1) % len(gas)
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
