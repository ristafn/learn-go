package main

import (
	"fmt"

	"github.com/ristafn/hello-world/library"
)

func main() {
	var student = library.Student{Name: "Ethan", Grade: 21}

	fmt.Println(student.Name)
	fmt.Println(student.Grade)
}
