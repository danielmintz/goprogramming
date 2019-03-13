package main

import "fmt"

func main() {

	fmt.Println(Mysum(2, 3))
	fmt.Println(Mysum(4, 5, 5, 3))
	fmt.Println(Mysum(2, 3, 6, 7, 8))
}

func Mysum(xs ...int) int {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return sum
}
