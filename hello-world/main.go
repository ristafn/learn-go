package main

import "fmt"

func main() {
  // var <variable-name> <type-data>
	var firstName string = "john"
	var lastName string
	lastName = "wick"

  // multi-variable
  // var first, second, third string = "satu", "dua", "tiga"

  // reserved variable menampung junk value
  _ = "belajar Golang"
  // name, _ := "john", "wick"

  // pointer
  // name := new(string)
  // fmt.Println(name)
  // fmt.Println(*name)


	fmt.Printf("hallo %s %s!\n", firstName, lastName)
}
