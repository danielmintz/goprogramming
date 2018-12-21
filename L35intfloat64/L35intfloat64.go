package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 42     // this is an int which is a whole number
	y := 42.345 // type here is a float62 for decimalsknown a sloating or real number

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(runtime.GOOS)   //operating systm of computer
	fmt.Println(runtime.GOARCH) // architecture of system

}

// nb uint would be a positive wherea int can be positive and negative
