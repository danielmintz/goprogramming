package main

import "fmt"

func main() {
	fmt.Println("1) hello evryone this is the best way to learn go")

	foo() // bit like an anchor and then returns
	fmt.Println("3) Something more ")

	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}

	}
}

func foo() {
	fmt.Println("2) Im in foo")

}

//Control Flow:
//Sequential
//Iterative / Loop
//Conditional
