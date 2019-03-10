package main

import (
	"fmt"
	"goprogramming/ToddGoCourse/L12Jedi/L12Jedi1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {

	dog1 := canine{
		name: "Molly",
		age:  dog.Years(11),
	}

	fmt.Println(dog1)

}
