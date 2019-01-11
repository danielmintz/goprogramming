package main

import (
	"fmt"
)

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Daniel",
		last:  "Mintz",
		age:   47,
	}

	fmt.Println(p1)

}
