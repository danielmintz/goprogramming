package main

import "fmt"

func main() {
	x := []int{5, 4, 6, 7, 8}
	fmt.Println(x)
	x = append(x, 12, 13, 14, 15)
	fmt.Println(x)
	y := []int{77, 88, 99, 101}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
}
