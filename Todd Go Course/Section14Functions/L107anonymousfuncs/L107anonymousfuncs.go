package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("hey there, this is an anonymous func")
	}()

	func(x int) {
		fmt.Println("hey there, this is an anonymous func pulling in a variable of type int:", x)
	}(42)

	func(y string) {
		fmt.Println("hey this is an anonymous func pulling in a variable of type string:", y)
	}("this is a string")

	func(z bool) {
		fmt.Println("hey this is an anonymous func pulling in a variable of type bool:", z)
	}(true)

	//fmt.Println("interesting")

}

//Anonymous self-executing func
