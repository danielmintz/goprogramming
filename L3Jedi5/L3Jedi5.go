package main

import (
	"fmt"
)

func main() {
	x := 10
	for {
		if x > 100 {
			break
		}
		fmt.Println(x % 4)
		x++
	}
}

//Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
