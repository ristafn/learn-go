package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
  fmt.Print("Type some number : ")
	fmt.Scan(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println(input, "is not number")
    fmt.Print(err.Error())
		return
	}

	fmt.Println(number, "is number")
}
