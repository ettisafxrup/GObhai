package main

import "fmt"

type User struct {
	Name string
	Role string
}

func printUser(user *User) {
	fmt.Println(user.Name)
}

func main() {
	fmt.Println("Hello World!")
	rup := User{"Rup", "Admin"}
	ruh := User{"Ruh", "Studnet"}

	printUser(&rup)
	printUser(&ruh)
}
