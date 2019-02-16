package main

import "fmt"

func main() {
	x := []int{5, 4, 6, 7, 8}
	fmt.Println(x)
	x = append(x, 12, 13, 14, 15)
	fmt.Println(x)
	y := []int{77, 88, 99, 101}1§
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
}

//To append values to a slice, we use the special built in function append. 
//This function returns a slice of the same type.
//note when curly {} and when brackets ()
//note ... (include everything) research variadic function.. an infinite number
