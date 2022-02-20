package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("Luas lingkaran\t\t: %.2f\n", area)
	fmt.Printf("Keliling lingkaran\t: %.2f\n", circumference)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

// predefined return value
func _(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}
