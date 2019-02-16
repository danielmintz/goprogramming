package main

import (
	"fmt"
)

var g func()

func main() {

	h := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

	}
	h()
	fmt.Println("this is done")

	g = h
	g()

	f := func(x int) {
		fmt.Println("awesome func as a variable being called and run", x)
	}
	f(190)

}

//Assign a func to a variable, then call that func
