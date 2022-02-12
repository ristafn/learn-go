package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"} // slice doens't have limit

	fmt.Println(fruits[0])

	fmt.Println()

	// var newFruits = fruits[0:2]
	// fmt.Println(newFruits)

	// explain type data reference
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	baFruits[0] = "pinnaple"

	fmt.Println()

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	fmt.Println()

	fmt.Println("Jumlah buah : ", len(fruits))
	fmt.Println("Jumlah buah : ", cap(fruits))

	aFruits = fruits[0:3]
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	bFruits = fruits[2:3]
	fmt.Println(cap(bFruits))

	var cFruits = append(fruits, "papaya")

	fmt.Println(cFruits)

	fmt.Println()

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n) // len(dst)
}
