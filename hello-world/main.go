package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println()

	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println()

	// infinit loop
	var j = 0

	for {
		fmt.Println("Angka", j)
		j++

		if j == 5 {
			break // stop the infiti loop when j equal 5
		}
	}

	// break and continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Println()

	// nested loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	fmt.Println()

	// label loop
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
