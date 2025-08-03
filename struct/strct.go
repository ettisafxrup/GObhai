package main

import "fmt"

type User struct {
	Name string
	Age  string
}

func (user User) greetUser() {
	fmt.Println("Hello", user.Name)
}

func main() {

	var user1 User = User{Name: "Rup", Age: "22"}

	user2 := User{Name: "Mahi", Age: "22"}

	fmt.Println(user1.Name, user1.Age)
	fmt.Println(user2.Name, user2.Age)

	user1.greetUser()
}
