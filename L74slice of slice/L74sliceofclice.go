package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5])
	fmt.Println(x[1:])
	fmt.Println(x[3:6])
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}
	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])
	}

}

//We can slice a slice which means that we can cut parts of the slice away. We do this with the colon operator.
