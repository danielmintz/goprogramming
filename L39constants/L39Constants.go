package main

import "fmt"

// constnt is written as const and means it stays the same always and cant be changed i.e. it is locked

func main() {
	//const a = 42
	//const b = 42.3334
	//const c = "Hello World"

	//alternative way of writing above
	const (
	//	a = 42
	//	b = 42.334
	//	c = "hello world"
	)

	//another alternative way of writing above
	const (
		a int     = 42
		b float64 = 42.334
		c string  = "hello world"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}

//Constant is a simple, unchanging value
//Only exist at compile time.
//there are TYPED and UNTYPED constants
//const hello = "Hello, World"
//const typedHello string = "Hello, World"
//UNTYPED constant
// a constant value that does not yet have a fixed type
//“constant of a kind”
//not yet forced to obey the strict rules that prevent combining differently typed values
//An untyped constant can be implicitly converted by the compiler.
//It is this notion of an untyped constant that makes it possible for us to use constants in Go with great freedom.
//This is useful, for instance
//what is the type of 42?
//int?
//uint?
//float64?
//if we didn’t have UNTYPED constants (constants of a kind), then we would have to do conversion on every literal value we used
//and that would suck
