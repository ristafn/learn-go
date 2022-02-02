package main

import "fmt"

func main() {
  const firstName string = "john"
  const _ = "wick" // type inference

  fmt.Print("halo ", firstName, "!\n") // Print separator without space
}
