package main

import "fmt"

func printArr(nums *[5]int) {
	fmt.Println(nums)
}

func main() {
	fmt.Println("Hello World!")
	nums := [5]int{1, 2, 3, 4, 5}
	printArr(&nums)
}
