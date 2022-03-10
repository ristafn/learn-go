package main

import "fmt"

func main() {
	var numberA = 4

	fmt.Println(numberA)

	change(&numberA, 10)

	fmt.Println(numberA)
}

func change(original *int, value int) {
	*original = value
}
