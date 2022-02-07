package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Steve"
	names[1] = "Tony"

	fmt.Println(names[0], names[1])

	var fruits = [2]string{"jeruk", "Tomat"}

	fmt.Println("Jumlah element : ", len(fruits))
	fmt.Println(fruits)

	fmt.Println()

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("Data array \t:", numbers)
	fmt.Println("Jumlah elemen \t:", len(numbers))

	fmt.Println()

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Elemen %d : %s\n", i, fruits[i])
	}

	fmt.Println()

	for i, fruit := range fruits {
		fmt.Printf("Elemen %d : %s\n", i, fruit)
	}

  fmt.Println()

  // if i not used
  for _, fruit := range fruits {
    fmt.Printf("Nama buah : %s\n", fruit)
  }

  fmt.Println()

  var fruits2 = make([]string, 2)
  fruits2[0] = "Apple"
  fruits2[1] = "Manggo"

  fmt.Println(fruits2)
}
