package main

import (
	"fmt"
)

func main() {
	for a := 65; a <= 90; a++ {
		fmt.Println(a)
		for b := 1; b <= 3; b++ {
			fmt.Printf("\t%#U\n", a)
		}

	}
}

//Print every rune code point of the uppercase alphabet three times. Your output should look like this:
//65
//U+0041 'A'
//U+0041 'A'
//U+0041 'A'
//66
//U+0042 'B'
//U+0042 'B'
//U+0042 'B'
//… through the rest of the alphabet characters
