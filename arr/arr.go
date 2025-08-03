package main

import "fmt"

var (
	myArr = [5]int{1, 2, 3, 4, 5}
)

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	fmt.Println(myArr)
}
