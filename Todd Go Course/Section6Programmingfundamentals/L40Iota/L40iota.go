package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e //ease of programming can drop iota after first one and Go works it out
	f
)

func main() {

	fmt.Println(a, b, c, d, e, f)
}

//Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.
// It is reset to 0 whenever the reserved word const appears in the source. It can be used to construct a set of
//related constants:
