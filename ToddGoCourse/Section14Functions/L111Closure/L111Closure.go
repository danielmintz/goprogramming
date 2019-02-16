package main

import (
	"fmt"
)

var x int

func main() {

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int { //from here is is enclosed i.e. closure. hence brings in addition of x each time. If above was belwo then alwats be 1
		x++
		return x
	}
}

//one scope enclosing other scopes
//variables declared in the outer scope are accessible in inner scopes
//closure helps us limit the scope of variables
