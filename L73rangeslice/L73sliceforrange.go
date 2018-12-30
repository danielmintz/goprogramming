package main

import (
	"fmt"
)

func main() {
	x := []int{5, 4, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		println(i, v)
	}

}

//We can loop over the values in a slice with the range clause. We can also access items in a slice by index position.
