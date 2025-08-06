package main

import "fmt"

func main() {
	x := []int{}
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5)

	fmt.Println(x)
	fmt.Println(y)
}
