package main

import (
	"fmt"
)

func main() {

	defer foo()
	fmt.Println("this is the start")

}

func foo() {
	defer func() {
		fmt.Println("This is a func defered in a func")
	}()
	fmt.Println("today this came second")

}
