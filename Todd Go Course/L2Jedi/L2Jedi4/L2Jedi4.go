package main

import (
	"fmt"
)

func main() {
	v := 42
	fmt.Printf("%d\t%b\t%#x", v, v, v)
	x := v << 1
	fmt.Printf("\n%d\t%b\t%#x", x, x, x)
}
