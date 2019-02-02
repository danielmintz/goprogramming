package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {

	p1 := person{
		name: "Daniel",
	}
	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	(*p).name = "Debra"
}
