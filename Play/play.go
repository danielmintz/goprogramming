package main

import "fmt"

func main() {

	n := foo()
	a, b := bar()
	fmt.Println(n, a, b)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 617, "yooo"
}

//func (r receiver) identifier(parameters) (returns) { code }

//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results
