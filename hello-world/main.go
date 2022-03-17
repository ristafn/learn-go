package main

import (
	f "fmt" // aliasing

	. "github.com/ristafn/hello-world/library" // make library same level with main.go
)

func main() {
	var student = Student{Name: "Ethan", Grade: 21}

	f.Println(student.Name)
	f.Println(student.Grade)
}
