package main

import (
	"fmt"
)

func main() {

	//Efficient way to stop new arrays being created in th ebackground = driving increased
	//run time by defining capacity upfront
	//make needs TYPE, LENGTH and CAPACITY

	x := make([]int, 10, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[4] = 17
	x[0] = 1
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4567)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4568, 4569, 5000, 5001)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// as at max capacity next one doubles array size automatically form 15 to 30. so increaing run // processing time.

	x = append(x, 6000)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

// MAKE:
//Slices are built on top of arrays. A slice is dynamic in that it will grow in size.
// The underlying array, however, does not grow in size. When we create a slice,
//we can use the built in function make to specify how large our slice should be and
//also how large the underlying array should be. This can enhance performance a little bit.
//make([]T, length, capacity)
//make([]int, 50, 100)
