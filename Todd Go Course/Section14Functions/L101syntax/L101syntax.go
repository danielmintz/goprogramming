package main

import (
	"fmt"
)

//we call our func and pass in arguments (if any)
func main() {
	foo()
	bar("James")
	s1 := woo("Miss MoneyPenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming") // 2 variables one is the string other the bool
	fmt.Println(x)
	fmt.Println(y)
}

//we define our func with parameters (if any)
func foo() {
	fmt.Println("Hey this is foo")
}

// the paremeter below will use the argument above in func main
func bar(s string) {
	fmt.Println("Hello", s)
}

// using return
func woo(st string) string {
	return fmt.Sprint("Hello woo, this is ", st)
}

//using multiple returns
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprintln(fn, ln, `, says "hello" `)
	b := true
	return a, b
}

//Syntax
//func (r receiver) identifier(parameters) (returns) { code }
//know the difference between parameters and arguments
//everthing in go in PASS BY VALUE
