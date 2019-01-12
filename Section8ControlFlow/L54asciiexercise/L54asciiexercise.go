package main

import (
	"fmt"
)

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\t%#x\n", i, i, i)

		//fmt.Printf("%q\t%s\t", i, i)

		//s := string(i)
		//fmt.Println(s)
		//fmt.Printf("%q\t%s\t", s, s)
	}

}

//Loop - printing ascii
//We can use format printing to print out the
//decimal value with %d
//hex value with %#x
//unicode code point with %#U
//a tab with \t
//a new line with \n
