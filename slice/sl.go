package main

import "fmt"

func changeSlice(s []int) []int {
	s[0] = 69
	return s
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	y := x[4:] // 5, 6, 7

	y[0] = 69
	fmt.Println(y[0])
	fmt.Println(x)

	// a := changeSlice(y) // 69, 6, 7

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(a)
}
