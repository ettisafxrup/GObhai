package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, Mr " + name)
}

func main() {
	var x string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&x)
	fmt.Println(greet(x))

}
