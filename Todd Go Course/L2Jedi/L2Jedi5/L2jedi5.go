package main

import (
	"fmt"
)

func main() {
	x := `Hello World here is  raw string
	literal "you see ANOTHER THING"`
	fmt.Println(x)
	fmt.Printf("%T", x)
}
