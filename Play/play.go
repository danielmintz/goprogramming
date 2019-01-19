package main

import "fmt"

func main() {
	defer foo()
	xi := []int{2, 3, 4, 5, 6, 7, 8}
	x := sum(xi...) // unfurling the slice
	fmt.Println("Hi this is the total", x)

}

func sum(x ...int) int {
	sum := 0

	for i, v := range x {
		fmt.Println(i, v, sum)
		sum += v
	}
	return sum

}
func foo() {
	fmt.Println("Hi")
}
