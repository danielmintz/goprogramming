package main

import "fmt"

func main() {
	x := 42 //:- this is the short decalrtion stating what variable i.e 42 is related ot an idenitier i.e. x
	fmt.Println(x)
	x = 90 // as variable x is alreayd declared we can use = to define .N
	fmt.Println(x)
	foo()
	y := 23 + 17 // 23 +7 is an expression and adding in the y bit is a statement. Multiple statements make a programm . NB the + is an operator and the numbers either side the operANDS
	fmt.Println(y)
}

func foo() {
	fmt.Println("little break and back")
}

// 6 statement is this program and each statement is made up of expressions.
