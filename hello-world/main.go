package main

import (
	"fmt"
)

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	newNumbers := func(min int) []int {
		var r []int // [3, 4, 3, 4, 3]

		for _, e := range numbers {
			if e < min { // e = 3, 3 < 3 false
				continue
			}

			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)
}
