package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Grade int
}

func (s *Student) SetName(name string) {
    s.Name = name
}

func main() {
	var s1 = &Student{Name: "john wick", Grade: 2}
    fmt.Println("nama :", s1.Name)

    var reflectValue = reflect.ValueOf(s1)
    var method = reflectValue.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("wick"),
    })

    fmt.Println("nama :", s1.Name)
}
