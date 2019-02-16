package main

import (
	"fmt"
)

func main() {

	if x := 42; x == 42 {
		fmt.Println(x)
	}
	//fmt.Println(x)
}

// initialisation stament (i.e. using a semi colon... this narrows scope as narrow as possible).
// So here the initilaisation statement is up to the smei colone i,e, if x  := 42;. this then shortens
// the scope to the first curly } so a command outside this curly relating to x no longer works (try uncommenting)
