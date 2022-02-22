package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f\n", avg)

	fmt.Println(msg)

	var hobbies = []string{"Fishing", "Diving", "Sleep"}

	yourHobbies("Budi", hobbies...)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))

	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is %s\n", name)
	fmt.Printf("My Hobie is %s\n", hobbiesAsString)
}
