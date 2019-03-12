package main

import "fmt"

func main() {
	fmt.Println(MySum(1, 2, 3, 4))
	fmt.Println(MySum(3, 6, 5, 6, 5, 7))
}

//Sum is cool
func MySum(xs ...int) int {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return sum
}
