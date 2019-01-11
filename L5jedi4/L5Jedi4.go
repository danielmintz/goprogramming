package main

import (
	"fmt"
)

func main() {

	p1 := struct {
		doors       int
		colour      string
		dashoptions []string
		trim        map[int]string
	}{
		doors:  6,
		colour: "red",
		dashoptions: []string{
			"Mahogany",
			"Oak",
			"Plastic",
		},
		trim: map[int]string{
			200: "basic range",
			300: "top range",
		},
	}

	for i, v := range p1.dashoptions {
		fmt.Println(i, v)
		for i1, v2 := range p1.trim {
			fmt.Println(i1, v2)
		}
	}

}
