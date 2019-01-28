package main

import (
	"fmt"
)

func main() {

	a := 42
	fmt.Println(a) // 42

	fmt.Println(&a) // 	0xc000088008 gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // '&' prints out the value stored at the address when you have the address

	b := &a
	fmt.Println(b) // 0xc000088008

	fmt.Println(*b) //42 '*' prints the value behind the location

	fmt.Printf("%T\n", b)

	fmt.Println(*&a)

	*b = 45
	fmt.Println(a)
	// new one

	c := 30
	d := &c
	*d = 35
	fmt.Println(c)

}

// the above code makes b a pointer to the memory address where an int is stored
// b is of type "int pointer"
// *int -- the * is part of the type -- b is of type *int

// b is an int pointer;
// b points to the memory address where an int is stored
// to see the value in that memory address, add a * in front of b
// this is known as dereferencing
// the * is an operator in this case

// this is useful
// we can pass a memory address instead of a bunch of values (we can pass a reference)
// and then we can still change the value of whatever is stored at that memory address
// this makes our programs more performant
// we don't have to pass around large amounts of data
// we only have to pass around addresses

// everything is PASS BY VALUE in go, btw
// when we pass a memory address, we are passing a value
