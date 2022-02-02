package main

import "fmt"

func main() {
  // numeric non-decimal
  var positiveNumber uint8 = 89
  var negativeNumber = -1243423644 // int32

  fmt.Println("Bilangan non-desimal")
  fmt.Printf("bilangan positif: %d\n", positiveNumber)
  fmt.Printf("bilangan negatif: %d\n", negativeNumber)
  fmt.Println("=====================================")

  // numeric decimal
  var decimalNumber = 2.62

  fmt.Println("Bilangan desimal")
  fmt.Printf("bilangan desimal: %f\n", decimalNumber) // float32
  fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) //float32
  fmt.Println("=====================================")

  // boolean
  var exist bool = true

  fmt.Println("Boolean")
  fmt.Printf("exist? %t \n", exist)
  fmt.Println("=====================================")

  // string
  var message string = "Hallo"
  var anotherMessage = `Nama saya "John Wick".
  Salam kenal.
  Mari belajar "Golang".`

  fmt.Println("String")
  fmt.Printf("message: %s\n", message)
  fmt.Println(anotherMessage)
  fmt.Println("=====================================")

  
}
