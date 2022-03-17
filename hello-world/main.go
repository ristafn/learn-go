package main

import "fmt"

type User struct {
	Name string
	Pet  []string
}

func (u *User) newPet() {
	fmt.Println("Get the old value", *u)
	u.Pet = append(u.Pet, "Lucy")
	fmt.Println("Get the new value", *u)
}

func main() {
	u := User{"Kevin", []string{"Cat"}}
	fmt.Println("get a old user", u)

	u2 := &u

	u2.newPet()
	fmt.Println("get a new user", u)
}
