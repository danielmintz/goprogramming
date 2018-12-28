package main

import (
	"fmt"
)

func main() {
	x := 45
	if x == 42 {
		fmt.Println("Our number was 42")
	} else if x == 41 {
		fmt.Println("our number was 41")
	} else if x == 43 {
		fmt.Println("our number was 43")
	} else if x == 44 {
		fmt.Println("our number was 44")
	} else {
		fmt.Println("our number was not 41, 42, 43 or 44")
	}
}
