package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&x)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&y)
	fmt.Println("Your Add Result:", add(x, y))
}
