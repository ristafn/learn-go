package library

import "fmt"

// Exported
func SayHello(name string) {
	fmt.Println("Hello")
  introduce(name) // make exported the unexported func
}

// Unexported
func introduce(name string) {
  fmt.Println("Nama saya", name)
}
