package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)                //three dots after xi unfurls to allow the int in from below.
	fmt.Println("the total is", x) // It takes a slice of variadic ""...int" below and lets it in. It
} //(i.e.the slice) has the same underlying array as variadic below.

func sum(x ...int) int { //variadic int must be last parameter. You could have a string
	fmt.Println(x) //before for eg but not after
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(len(x))
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("the value at position", i, "which is", v, "the total is now", sum)
	}
	return sum
}

//When you have a slice of some type, you can pass in the individual values in a slice by using the “...” operator.
