package main

import "fmt"

func main() {
	// lets talk about the structs
	// and lets talk about the diff things that change with structs that are varying with this currently
	newuser := User{"Padam", "user@gmail.com", 28, true}
	Greet(newuser)
	AgeIncrement(&newuser)
	fmt.Printf("New age %v", newuser.Age)
}

func Greet(u User) {
	fmt.Printf("hello my name is %v\n", u.Name)
}

func AgeIncrement(u *User) {
	u.Age++
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
