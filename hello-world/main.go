package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("name cannot be empty")
	}

	return true, nil
}

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	if valid, error := validate(name); valid {
		fmt.Println("hello", name)
	} else {
		fmt.Println(error.Error())
	}
}
