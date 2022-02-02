package main

import "fmt"

func main() {
  var value uint8 = (((2 + 6) % 3) * 4 - 2) / 3
  var isEqual = (value == 2)

  fmt.Printf("nilai %d (%t) \n", value, isEqual)
  fmt.Println("===============================")


  var left bool = false
  var right bool = true

  fmt.Println("Logic operator")
  var leftAndRight bool = left && right
  fmt.Printf("left && right \t(%t) \n", leftAndRight)

  var leftOrRight bool = left || right
  fmt.Printf("left || right \t(%t) \n", leftOrRight)

  var leftReverse bool = !left
  fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
