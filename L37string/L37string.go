package main

import (
	"fmt"
)

func main() {
	s := `"Hello world"
	:)
	` // string works with quotes or backticks (latter include new lines etc)
	// a string is made up of a slice of bytes...nb each letter has a numeric value
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // convert s to slice of byte which  NB byte is same as int8
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

}
