package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []struct{
		person
		_ int
	}{
		{person: person{name: "Josee", age: 20}},
		{person: person{name: "Steve", age: 23}},
		{person: person{name: "Kepin", age: 19}},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}
