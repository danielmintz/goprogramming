package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // Print type
	fmt.Printf("%b\n", y) // Print binary
	fmt.Printf("%x\n", y) // Print hexadecimal
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("\n%#x\n%b\n%x", y, y, y) // n = new line
	fmt.Printf("\t%#x\t%b\t%x", y, y, y) // t = tab
	s := fmt.Sprintf("\t%#x\t%b\t%x", y, y, y)
	fmt.Println(s)

	fmt.Printf("%v", y) //%v is the default value

}
