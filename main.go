package main

import (
	"fmt"

	"rupMaths/mathlib"
)

func init() {
	fmt.Println("Welcome to GObhai")
}

// main function that asks user for two numbers and prints their sum
func main() {

	var x int
	var y int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&x)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&y)

	func(a, b int) {
		fmt.Println("Joy Bangla")
		fmt.Println(a + b)
	}(x, y)

	fmt.Println("Your Add Result from mathlib Package:", mathlib.Add(x, y))
}
