package main

import (
	"fmt"
)

func main() {
	//fmt.Println(bar()())    //replicates below with simpler code

	x := bar()
	i := x()
	fmt.Printf("%T\n", x)
	fmt.Println(i)

	//fmt.Println(x())

	//fmt.Println(bar()())

}

func bar() func() int {
	return func() int {
		return 451
	}
}

//You can return a func from a func.
