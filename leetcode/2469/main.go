package main

import "fmt"

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}

func main() {
	celsius := 10.00
	fmt.Println(convertTemperature(celsius))
}
