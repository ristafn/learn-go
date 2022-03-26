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

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	defer catch()

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	if valid, error := validate(name); valid {
		fmt.Println("hello", name)
	} else {
		panic(error.Error())
		fmt.Println("end")
	}
}
