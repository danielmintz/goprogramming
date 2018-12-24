package main

import (
	"fmt"
)

func main() {
	s := "H"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	//convert the string "H" to slice of byte
	bs := []byte(s)
	fmt.Println(bs)

	// assign to index position 0 i.e. just number 72
	n := bs[0]
	fmt.Println(n)

	// Print our in decimal, binary and hexadecimal to show the different numeral values of 72.
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%x\n", n)

}
